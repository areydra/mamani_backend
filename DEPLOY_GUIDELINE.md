1. Open cpanel file manager -> upload mamani-backend into folder mamani_backend
2. Change permission file mamani-backend. These are permission you need to check:
    - User: Read, Write, and Execute
    - Group: Read
    - World: Read
3. Open putty and login ssh
4. There are several command to run the service:
    - Run => ./mamani-backend
    - Run background => nohup ./mamani-backend
    - Check background running => ps aux | grep mamani-backend
    - Stop background running => kill -9 process_id
