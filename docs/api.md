## Lunch Order REST API

#### /auth/google
- GET: Verify access token
    - Request:
        ```json
        {
            "email" : "token",
            "token" : "string"
        }
        ```
    - Response:
        - status code: 200(OK) / 401(Unauthorized)

#### /auth/google/logout
- GET: Logout 
    - Request:
        ```json
        {
            "email" : "token",
            "token" : "string"
        }
        ```
    - Response:
        - status code: 200

#### /auth/google/login-url

- GET: Generate Google Login URL
    - Request:
        ```json
        {}
        ```
    - Response:
        ```json
        {
            "redirect_url" : "string",
            "token" : "string",
            "client_id" : "string",
        }
        ```

#### /auth/google/callback

- GET: Verify user login
    - Request:
        ```json
        {
            "email" : "string",
            "token" : "string",
        }
        ```
    - Response:
        ```json
        {
            "id" : 0,
            "name" : "string",
            "email" : "string",
            "token" : "string"
        }
        ```

#### /menus

- GET: Get latest menu
    - Request:
        ```json
        {}
        ```
    - Response:
        ```json
        {   "menu" : {
                "id" : 0,
                "owner_id" : 0,
                "name" : "string",
                "deadline" : "string",
                "payment_reminder" : "string",
                "create_at" : "string",
                "status" : 0,
            },
            "items" : [
                {
                    "id" : 0,
                    "item_name" : "string",
                    "users" : [
                        {
                            "id" : 0,
                            "user_name" : "string"
                        }
                    ]
                }
            ],
            "people_in_charge" : [
                {
                    "user_id" : 0,
                    "user_name" : "string"
                }
            ]
        }
        ```

- POST: Create menu
    - Request:
        ```json
        {   "menu" : {
                "owner_id" : 0,
                "name" : "string",
                "deadline" : "string",
                "payment_reminder" : "string"
            },
            "item_names" : [],
        }
        ```
    - Response:
        ```json
        {   "menu" : {
                "id" : 0,
                "owner_id" : 0,
                "name" : "string",
                "deadline" : "string",
                "payment_reminder" : "string",
                "create_at" : "string",
                "status" :"string"
            },
            "items" : [
                {
                    "id" : 0,
                    "item_name" : "string"
                }
            ],
        }
        ```

#### /menus/{MenuID}/time

- POST: Modify menu's deadline and payment time
    - Request:
        ```json
            {
                "deadline" : "string",
                "payment_reminder" : "string",
            }
        ```
    - Response:
        ```json
        {   
            "deadline" : "string",
            "payment_reminder" : "string",
        }
        ```

#### /menus/{MenuID}/items

- POST: Add an items to menu
    - Request:
        ```json
        {
            "item_name" : "string"
        }
        ```
    - Response:
        ```json
        {   
            "item" : {
                "id" : 0,
                "item_name" : "string",
                "menu_id" : "string"
            },
        }
        ```

#### /menus/{MenuID}/items/{ItemID}

- DELETE: Delete an item from menu
    - Request:
        ```json
        {}
        ```
    - Response:
        ```json
        {   
            "item" : {
                "id" : 0,
                "item_name" : "string",
                "menu_id" : "string"
            },
        }
        ```

#### /menus/{MenuID}/users/{UserID}/orders

- GET: Get orders of an user
    - Request:
        ```json
        {}
        ```
    - Response:
        ```json
        {
            "items" : [
                {
                    "id" : 0,
                    "name" : "string"
                }
            ]
                  
        }
        ```
- POST: Create/Modify orders
    - Request:
        ```json
        {
            "item_ids" : []
        }
        ```
    - Response:
        ```json
        {
            "items" : [
                {
                    "id" : 0,
                    "name" : "string"
                }
            ]
        }
        ```
- DELETE: Cancel all orders of user
    - Request:
        ```json
        {}
        ```
    - Response:
        ```json
        {
            "items" : [
                {
                    "id" : 0,
                    "name" : "string"
                }
            ]
                  
        }
        ```
#### /menus/{MenuID}/people-in-charge

- GET: Get people in charge of the menu
    - Request:
        ```json
        {}
        ```
    - Response:
        ```json
        {
            "users" : [
                {
                    "id" : 0,
                    "name" : "string"
                }
            ], 
        }
        ```
- POST: Add people in charge
    - Request:
        ```json
        {
            "id" : []
        }
        ```
    - Response:
        ```json
        {
            "users" : [
                {
                    "id" : 0,
                    "name" : "string"
                }
            ],    
        }
        ```

### Error response
```json
{
    "error" : {
        "code" : "string",
        "message" : "string"
    }
}
```