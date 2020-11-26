echo ""
echo "CLEANING CONTAINER qa_server"
sudo docker container rm qa_server --force
echo "CLEANING IMAGE qa_server"
sudo docker image rm qa_server --force
echo ""
echo "BUILDING qa_server"
sudo docker build -t qa_server -f Dockerfile .
echo ""
echo "RUNNING qa_server"
# https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach
sudo docker run --network="host" --name qa_server -p 8080:8080 qa_server
