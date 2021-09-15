![license](https://img.shields.io/github/license/symmetric-project/node-backend)
![coverage](https://img.shields.io/github/license/symmetric-project/node-backend)
![stars](https://img.shields.io/github/stars/symmetric-project/node-backend)
![forks](https://img.shields.io/github/forks/symmetric-project/node-backend)
![issues](https://img.shields.io/github/issues/symmetric-project/node-backend)

# node-backend
The backend for a Symmetric node.

## Contributing
We have a contributing guide over [here](CONTRIBUTING.org).

## Community
You can contact us by joining our discord [here](https://discord.com/invite/AkSbP5JF).
## To run the backend server you need to:

### Create a PostgreSQL database with name of ```symmetric``` as well as set a password for the ```postgres``` user

### Create an ```.env``` file which holds sensitive data like passwords, instructions for Linux (Ubuntu, macOS and more).
Create file by running this command:
```
touch .env
```
Edit the file with any text editor and paste this:
```
MODE=dev
DATABASE_URL=postgres://postgres:REPLACE_THIS_WITH_THE_PASSWORD_OF_THE_LOCAL_SYMMETRIC_DATABASE@localhost:5432/symmetric
JWT_SECRET=REPLACE_THIS_WITH_A_RANDOMLY_GENEREATED_SHA256_STRING
COOKIE_DOMAIN_DEV=symmetric.localhost
COOKIE_DOMAIN_PROD=symmetric.REPLACE_THIS_WITH_THE_NODE_NAME.com
```
A way to generate a ```JWT_SECRET``` is to run the following command:
```
openssl rand -base64 64
```
After running it, you should see get a random string which you can use as a ```JWT_SECRET```

After this, run `gqlgen generate``` in the root folder. This will generate Go types from the `schema.graphqls` file.
