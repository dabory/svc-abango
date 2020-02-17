# gRPC KangAn How to

## Running the server and client

Note : Do not delete runner.conf !
When you copy the folder to the other folder, fresh command in docker-compose would not work
because root folder of the fresh is not written in docker by default. So you need to run 

`$ fresh -c runner.conf

`$ dkcu or docker-compose up -d

then it is to be affected by changing codes in the folder for compiling go lang codes automatically

!! (1) change runner.conf -> (2) run fresh successfully -> (3) fresh working properly.