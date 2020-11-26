// https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/
// https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver

// qa_db is the package that contains functions to connect Golang with MongoDB.
package qa_db

import (
    "fmt"
    "log"
    "strconv"
    "context"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// QA is a structure to store Questions with their Answers.
type QA struct {
    Id string `bson:"id"` // Id is a string that represents the QA.
    Question string `bson:"question"` // Question is a string that contains the question.
    QUser string `bson:"quser"` // QUser is the user who asked the question.
    Answer string `bson:"answer"` // Answer is a string that contains the answer to the question.
    AUser string `bson:"auser"` // AUser is the user who answered the question.
}

// QuestionId is an int that stores the next ID to be used by the struct QA.
var QuestionId int = 0

// collection is a MongoDB collection.
var collection *mongo.Collection

// ctx is a context variable to handle errors.
var ctx = context.TODO()

// MongoURI is the URI to connect to MongoDB.
var MongoURI = "mongodb://127.0.0.1:27017/"

// InitializeMongoDB initializes the connection to MongoDB.
func InitializeMongoDB() {
    fmt.Printf("Initializing MongoDB at %v\n", MongoURI)
    clientOptions := options.Client().ApplyURI(MongoURI)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database("qa").Collection("qas")
    fmt.Println("MongoDB initialized.")
    UpdateNewQuestionId()
    //ShowAllQuestions()
}

// UpdateNewQuestionId determines the QuestionId according to the number of QA records.
func UpdateNewQuestionId() {
    qas, err := GetAllQA()
    if err == nil {
        QuestionId = len(qas)
        fmt.Printf("QuestionId=%v\n", QuestionId)
    } else {
        fmt.Println(err)
    }
}

// ShowAllQuestions shows all the questions stored in MongoDB. It is a function for debugging.
func ShowAllQuestions() {
    qas, err := GetAllQA()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("All questions:")
    for _, qa := range qas {
        fmt.Printf("%+v\n", qa)
    }
    qa, _ := GetQA("1")
    fmt.Printf("GetQA(1): %+v\n", qa)  
}

// GetAllQA gets all the QA records.
func GetAllQA() ([]QA, error) {
    filter := bson.D{{}}
    return FilterQAs(filter)
}

// GetQA gets the first QA record with the corresponding ID.
func GetQA(id string) (QA, error) {
    filter := bson.D{primitive.E{Key: "id", Value: id}}
    qas, err := FilterQAs(filter)
    if err != nil {
        return QA {}, err
    } else {
        if len(qas) == 0 {
            return QA {}, fmt.Errorf("Question id %v not found.", id)
        } else {
            return qas[0], nil
        }
    }
}

// GetQuestionsOfUser gets all the questions asked by a user.
func GetQuestionsOfUser(quser string) ([]QA, error) {
    filter := bson.D{primitive.E{Key: "quser", Value: quser}}
    return FilterQAs(filter)
}

// GetAnswersOfUser gets all the answers of a user.
func GetAnswersOfUser(auser string) ([]QA, error) {
    filter := bson.D{primitive.E{Key: "auser", Value: auser}}
    return FilterQAs(filter)
}

// FilterQAs is a helper function to filter QA records.
func FilterQAs(filter interface{}) ([]QA, error) {
    var qas []QA
    cur, err := collection.Find(ctx, filter)
    if err != nil {
        return qas, err
    }
    for cur.Next(ctx) {
        var qa QA
        err := cur.Decode(&qa)
        if err != nil {
            return qas, err
        }
        qas = append(qas, qa)
    }
    if err := cur.Err(); err != nil {
        return qas, err
    }
    cur.Close(ctx)
    return qas, nil
}

// CreateQA creates a new QA record. It automatically handles the question ID.
func CreateQA(qa QA) (QA, error) {
    QuestionId += 1
    qa.Id = strconv.Itoa(QuestionId)
    _, err := collection.InsertOne(ctx, qa)
    return qa, err
}

// UpdateQA updates an existing question. If the question does not exist, an error is thrown.
func UpdateQA(qa QA) error {
    _, err0 := GetQA(qa.Id)
    if err0 != nil {
        return err0
    }
    filter := bson.D{primitive.E{Key: "id", Value: qa.Id}}
    update := bson.D{primitive.E{Key: "$set", Value: bson.D{
        primitive.E{Key: "question", Value: qa.Question},
        primitive.E{Key: "quser", Value: qa.QUser},
        primitive.E{Key: "answer", Value: qa.Answer},
        primitive.E{Key: "auser", Value: qa.AUser},
    }}}
    _, err := collection.UpdateOne(ctx, filter, update)
    return err
}

// DeleteQA deletes an existing question. If the question does not exist, an error is thrown.
func DeleteQA(Id string) error {
    filter := bson.D{primitive.E{Key: "id", Value: Id}}
    res, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }
    if res.DeletedCount == 0 {
        return errors.New("No questions were deleted")
    }
    return nil
}

// DeleteAllQAs deletes all the questions in the database.
func DeleteAllQAs() error {
    filter := bson.D{{}}
    _, err := collection.DeleteMany(ctx, filter)
    if err != nil {
        return err
    }
    QuestionId = 0
    return nil
}
