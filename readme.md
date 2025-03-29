
# Building a Backend Server with Golang and the Gin Framework

This project implements a backend server using the gin web framework in Golang. It includes logging, authentication middleware, and deployment as a daemon on an EC2 instance.

### Salient Features:

● The entrypoint for the code is the main package in the main.go file.

● The ‘gin’ web framework of golang was employed for the backend.

● The logs are being written to a temp/log/monitoring-api.log file for future inspection.

● Authentication middleware is implemented to check the API_KEY passed in the request. This is present in the middleware/auth.go file in middleware folder

● Four different routes for getting CPU, Memory, Disk and Bandwidth have been
implemented. The logic for this is written in the handlers/handler.go file.

● The server is listening on port 8080.

● A binary executable was created from this, which was then deployed to the EC2 and run
from as a daemon with systemctl commands.


### Steps for replicating this are as follows:
1. Download and unzip the monitoring-api folder.
2. In order to run the code locally for testing, download Go on your system.
3. Install dependencies with command: go mod tidy.
4. Create a binary file with the command: go build -o monitoring-api
5. Run the server with: ./monitoring-api
6. Test this locally by running: curl -H "API-KEY: cs218secret"
http://localhost:8080/cpu
7. If this is running well, we now start moving to EC2
8. The binary file created might not be appropriate for the EC2 instance so create one for
linux architecture with command: GOOS=linux GOARCH=amd64 go build -o
monitoring-api
9. Create a new EC2 instance of type t2.micro with Amazon Linux OS. Also create an
RSA key-pair for access (or use a previous one). In this case the name of the file is
cs218demo.pem and it is stored in the parent folder outside monitoring-api.
10. In order to allow for traffic from the internet on port 8080, add a new inbound rule to the
EC2 security group.
Click on EC2 instance ID >> Security Tab >> Click on Security Group >> Edit Inbound
Rule >> Add Rule >> Type: Custom TCP, Port: 8080, Source: 0.0.0.0/0 >> Save Rule
11. Move the binary file from the local to EC2 with the command:
scp -i ../../cs218demo.pem monitoring-api
ec2-user@ec2-54-81-116-95.compute-1.amazonaws.com:/home/ec2-user/
Replace the path to pem file with your path and the
hostname(ec2-user@ec2-54-81-116-95.compute-1.amazonaws.com) with the Public
IPv4 DNS of your EC2.
12. Connect to the EC2 with SSH: ssh -i "
../../cs218demo.pem"
ec2-user@ec2-54-81-116-95.compute-1.amazonaws.com
The monitoring-api file should be visible by running the ls command inside EC2. Replace
the hostname with yours.
13. Next create the folder for log file and give its access to ec2-user:
mkdir -p /home/ec2-user/temp/log/
chmod -R 755 /home/ec2-user/temp/log/
chown -R ec2-user:ec2-user /home/ec2-user/temp/log/
14. Now you can try running the program manually in EC2 with command:
./monitoring-api
15. To test the endpoint, call it from outside the EC2 (in a new terminal) with the command:
curl -H "API-KEY: cs218secret"
https://ec2-54-81-116-95.compute-1.amazonaws.com:8080/cpu
Replace the hostname with yours.
16. Next to create a daemon from this, create a new service file with command:
sudo nano /etc/systemd/system/monitoring-api.service
17. Add the following details to the file:

[Unit]
Description=Monitoring API Service
After=network.target

[Service]
ExecStart=/home/ec2-user/monitoring-api
Restart=always
User=ec2-user
Environment=PATH=/usr/local/bin:/usr/bin:/bin
Environment=HOME=/home/ec2-user
WorkingDirectory=/home/ec2-user

[Install]
WantedBy=multi-user.target

18. Reload the daemons: sudo systemctl daemon-reload
19. Enable the new daemon created: sudo systemctl enable monitoring-api
20. Start the monitoring-api daemon: sudo systemctl start monitoring-api
21. Check the status with the command: sudo systemctl status monitoring-api
The following should be its output
22. Test the endpoint again to verify.
23. Its logs can be viewed by looking at the log file: sudo nano
/home/ec2-user/temp/log/monitoring-api.log
24. Here is an example of some logs
25. Now the EC2 instance can be also stopped and restarted. The program will start running
on restart.
