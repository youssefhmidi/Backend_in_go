## Directories Objectifs 

### Controller directory "/controller/"
Contains all the functions logic that will connect with the models and the database 

- Example:
    - user requested a Login session, the controller will grab the response ,that he should give to the user, from the Model and send it back.

### Middleware directory "/middleware/"
Contains the functionnalities that the backend API will use to identify requesters

- Example:
    - check if the access token is validated or not 
### Routes directory "/routes/"
Contais all the possible routes that users can request to

### Test Directory "/test/"
All the testing files of the api functionnalities and for debuging

### Database Directory "/database/"
the database connector and initializer and has the repository functionalities

### Models Directory "/models/"
contains informations about the models and the response/request structure

### Bootstrap directory "/bootstrap/"
initializing the Env variables and the app

### cmd directory "/cmd/"
the main file

