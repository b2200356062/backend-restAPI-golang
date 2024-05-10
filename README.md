
## Melih Ercan

I haven't made a project with GO before so this is my first time.
I mainly used Gin Framework and Gorm library for this project. I tried to stick to MVC architecture pattern with controller and models. 
I implemented an extra signup function to create users in to the database.

## Functionalities:
    -Two types of users exist, "first" and "second" type users. First type can only work on their list
    and messages but second type users can see other user's lists and messages.

    -Users need to login before creating a TO-DO list and need to create a TO-DO list before creating 
    a message in them. If not logged in, system responds with http 401 unauthorized.

    -Each user has a limit of 1 list for TO-Do list.

    -I purposefully hid created_at, updated_at and deleted_at timestamps in the response from user.

    -As stated in the document, delete won't actually deletes but just modify the timestamps for updated_at
    and deleted_at.

    -2 types of users already exist in the database;
            second type user credentials:
                    melih2@hotmail.com
                    123123123
            first type user credentials:
                    melih@hotmail.com
                    123123123
    Can test the system with these users, also can create another users with name, email, password and user type.

    -System has response messages for list and message creation, update, and deletion.

    -Users just enter list completion rate when creating a list and enter content and completion status for
    creating the messages. Other things created automatically.



