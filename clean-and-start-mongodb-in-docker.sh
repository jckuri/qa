echo ""
echo "CLEANING CONTAINER qa_mongo"
sudo docker container rm qa_mongo --force
echo "CLEANING IMAGE mongo"
sudo docker image rm mongo --force
echo ""
echo "RUNNING mongo"
sudo docker run --network="host" --name qa_mongo -p 127.0.0.1:27017:27017 mongo
