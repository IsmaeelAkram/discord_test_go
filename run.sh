docker build -t discord_test .
docker run -d --restart="always" --name discord_test -it discord_test