# Setup Instructions
To run the API locally, the local machine must have Docker Engine installed.

Then, create an `.env` file with the keys provided in the `env.example` file.

Enter the root directory. From the terminal, run the following command.
```
docker-compose up --build -d
```

After running the above command, the API will work locally and requests can be made.

 # API Testing dan Documentation
 The documentation of the API endpoints can be seen in the following Postman documenter in this link: https://documenter.getpostman.com/view/12531688/2s935ms55i

The automated testing suite of Postman can be accessed here: https://www.postman.com/richard-here/workspace/shopping-cart-api/collection/12531688-ffc54ac4-06ba-415b-9d75-492c3dd47f41?action=share&creator=12531688.

**Important note**: the Postman API workspace above contains pre-request and test scripts that configure the values to be used in the requests for automated testing. To achieve expected behavior, **do not run the endpoints manually**. Instead, run it using automatically by using the following steps.
- Ensure that the local webserver app is already running by running the Docker containers
- Set the environment in the workspace to `Soluix API Environment`
- Right click on the collection `Soluix Automated API Test Suites`
- Click on `Run Collection`
- Make sure the value of `Iterations` is `1` and `Delay` is `50`
- Click on the run button
- See the test result