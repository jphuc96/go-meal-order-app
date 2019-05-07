## Lunch Order REST API

#### /auth/google/logout
- POST: Logout 
    - Header: 
        - access_token: [base64_token]
    - Request:
        ```json
        {}
        ```
    - Response:
        - status code: 200

#### /auth/google/callback

- GET: Handle oauth2 callback
    - Header: 
        - client_id: 
        - id_token: 
    - Request:
        ```json
        {}
        ```
    - Response:
        ```json
        {   
            "id" : 0,
            "name" : "string",
            "email" : "example@gmail.com",
            "access_token" : "string",
        }
        ```

#### /menus

- GET: Get latest menu
    - Header: 
        - access_token: [base64_token]
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
    - Header: 
        - access_token: [base64_token]
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
    - Header: 
        - access_token: [base64_token]
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
    - Header: 
        - access_token: [base64_token]
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

#### /items/{ItemID}

- DELETE: Delete an item from menu
    - Header: 
        - access_token: [base64_token]
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
    - Header: 
        - access_token: [base64_token]
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
    - Header: 
        - access_token: [base64_token]
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
    - Header: 
        - access_token: [base64_token]
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
    - Header: 
        - access_token: [base64_token]
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

### Error response
```json
{
    "error" : {
        "code" : "string",
        "message" : "string"
    }
}
```

### After deadline, below APIs will be disable:
 - POST /menus/{MenuID}/users/{UserID}/orders
 - DELETE /menus/{MenuID}/users/{UserID}/orders