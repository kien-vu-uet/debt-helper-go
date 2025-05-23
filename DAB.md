# Design APIs

## Table of Contents
1. [Common](#common)
2. [Login](#login)
3. [User](#user)
4. [Groups](#groups)
5. [Group Transactions](#group-transactions)
6. [Debt Calculation](#debt-calculation)
## Common: 
*Prefix*: `/api/v1`

## Default: 
- `/health`: 
    - Method: `GET`
    - Description: Health check endpoint to verify if the API is running.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "status": "ok"
        }
        ```
- `/info`:
    - Method: `GET`
    - Description: Provides information about the API, such as version and documentation.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {`
            "version": "1.0.0",
            "documentation": "https://example.com/docs"
        }
        ```

---

## Login:
*Prefix*: `/login`

- `/access-token`:
    - Method: `POST`
    - Description: Authenticates a user and returns a token.
    - Request Body: 
    ```json
    {
        "username": "string",
        "password": "string"
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "token": "string"
        }
        ```
- `/access-token/extend`:
    - Method: `POST`
    - Description: Extends the validity of an existing access token.
    - Request Body: 
    ```json
    {
        "token": "string"
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "token": "string"
        }
        ```
- `/refresh-token/revoke`:
    - Method: `POST`
    - Description: Refreshes the access token using a refresh token.
    - Request Body: 
    ```json
    {
        "refresh_token": "string"
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "message": "Refresh token revoked successfully"
        }
        ```
- `/verify-username`:
    - Method: `POST`
    - Description: Verifies if a username is available.
    - Request Body: 
    ```json
    {
        "username": "string"
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "available": true
        }
        ```

---

## User:

*Prefix*: `/users`

- `/signup`:
    - Method: `POST`
    - Description: Registers a new user.
    - Request Body: 
    ```json
    {
        "username": "string",
        "password": "string",
        "email": "string"
    }
    ```
    - Response: 
        - Status: `201 Created`
        - Body: 
        ```json
        {
            "message": "User created successfully"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Username already exists"
        }
        ```
- `/me`:
    - Method: `GET`
    - Description: Retrieves the current user's profile information.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "username": "string",
                "email": "string",
                ...
            }
        }
        ```

- `/reset-password`:
    - Method: `POST`
    - Description: Resets the password for the current user.
    - Query Parameters:
        - `old_password`: The current password of the user.
        - `new_password`: The new password for the user.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "message": "Password reset successfully"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Old password is incorrect"
        }
        ```
- `/fullname`:
    - Method: `PATCH`
    - Description: Updates the current user's profile information.
    - Query Parameters:
        - `fullname`: The new full name of the user.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "message": "User updated successfully"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid fullname format"
        }
        ```
- `/avatar`:
    - Method: `PATCH`
    - Description: Updates the current user's avatar.
    - Multipart Form Data: (1 of the following)
        - `avatar`: The new avatar image file.
        - `avatar_url`: The URL of the new avatar image. (url or base64)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "message": "Avatar updated successfully"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid image format"
        }
        ```
- `/`:
    - Method: `GET`
    - Description: Retrieves a list of all users.
    - Query Parameters:
        - `page_size`: The page number for pagination.
        - `page_number`: The number of users per page.  (1 - 200)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data" {
                "page_size": 10,
                "page_number": 1,
                "total_pages": 100,
                "prev": 1 or null,
                "next": 3 or null,
                "data": [
                    {
                        "username": "string",
                        "email": "string"
                        ...
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid page size or number"
        }
        ```
- `/`:
    - Method: `POST`
    - Description: Creates a new user.
    - Request Body: 
    ```json
    {
        "username": "string",
        "email": "string"
        ...
    }
    ```
    - Response: 
        - Status: `201 Created`
        - Body: 
        ```json
        {
            "message": "User created successfully"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid email format"
        }
        ```

- `/{user_id}`
    - Method: `GET`
    - Description: Retrieves a specific user's profile information.
    - Path Parameters:
        - `user_id`: The ID of the user to retrieve.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "username": "string",
                "email": "string"
                ...
            }
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "User not found"
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
- `/{user_id}/`
    - Method: `DELETE`
    - Description: Deletes a specific user's profile information.
    - Path Parameters:
        - `user_id`: The ID of the user to delete.
    - Response: 
        - Status: `204 No Content`
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid email format"
        }
        ```
    - Error Response:   
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "User not found"
        }
        ```

---

## Groups:

*Prefix*: `/groups`

- *Common response* 
    - `200 OK`
    - `201 Created`
    - `204 No Content`
    - `400 Bad Request`
    - `401 Unauthorized`
    - `403 Forbidden`
    - `404 Not Found`

- `/`:
    - Method: `GET`
    - Description: Retrieves a list of all groups.
    - Query Parameters:
        - `page_size`: The page number for pagination.
        - `page_number`: The number of groups per page.  (1 - 200)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data" {
                "page_size": 10,
                "page_number": 1,
                "total_pages": 100,
                "prev": 1 or null,
                "next": 3 or null,
                "data": [
                    {
                        "group_id": "string",
                        "name": "string",
                        ...
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid page size or number"
        }
        ```
- `/`:
    - Method: `POST`
    - Description: Creates a new group.
    - Request Body: 
    ```json
    {
        "name": "string",
        ...
        "members": [
            "user_id_1",
            "user_id_2",
            ...
        ],
    }
    ```
    - Response: 
        - Status: `201 Created`
        - Body: 
        ```json
        {
            "data": {
                "group_id": "string",
                "name": "string",
                "description": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid group name"
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
- `/{group_id}`
    - Method: `GET`
    - Description: Retrieves a specific group's information.
    - Path Parameters:
        - `group_id`: The ID of the group to retrieve.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "group_id": "string",
                "name": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Group not found"
        }
        ```
- `/{group_id}`
    - Method: `PUT`
    - Description: Updates a specific group's information.
    - Path Parameters:
        - `group_id`: The ID of the group to update.
    - Request Body:
    ```json
    {
        "name": "string",
        ...
        "members": [
            "user_id_1",
            "user_id_2",
            ...
        ],
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "group_id": "string",
                "name": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{group_id}/name`:
    - Method: `PATCH`
    - Description: Partially updates a specific group's information.
    - Path Parameters:
        - `group_id`: The ID of the group to update.
    - Query Parameters:
        - `name`: The new name of the group.
    - Response:
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "group_id": "string",
                "name": "string",
                "description": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{group_id}/description`:
    - Method: `PATCH`
    - Description: Partially updates a specific group's description.
    - Path Parameters:
        - `group_id`: The ID of the group to update.
    - Query Parameters:
        - `description`: The new description of the group.
    - Response:
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "group_id": "string",
                "name": "string",
                "description": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{group_id}/avatar`:
    - Method: `PATCH`
    - Description: Updates a specific group's avatar.
    - Path Parameters:
        - `group_id`: The ID of the group to update.
    - Multipart Form Data: (1 of the following)
        - `avatar`: The new avatar image file.
        - `avatar_url`: The URL of the new avatar image. (url or base64)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "group_id": "string",
                "name": "string",
                "description": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{group_id}`
    - Method: `DELETE`
    - Description: Deletes a specific group.
    - Path Parameters:
        - `group_id`: The ID of the group to delete.
    - Response: 
        - Status: `204 No Content`
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{group_id}/settle`:
    - Method: `POST`
    - Description: Settles a specific group.
    - Path Parameters:
        - `group_id`: The ID of the group to settle.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
            {
            "data": {
                "group_id": "string",
                "name": "string",
                "description": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```


---

- `/{group_id}/members/`:
    - Method: `POST`
    - Description: Adds some users to a specific group.
    - Path Parameters:
        - `group_id`: The ID of the group to add the users to.
    - Request Body:
    ```json
    {
        "user_id": [
            "user_id_1",
            "user_id_2",
            ...
        ],
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "group_id": "string",
                "name": "string",
                "description": "string",
                ...
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{group_id}/members/{member_id}`:
    - Method: `DELETE`
    - Description: Removes a specific user from a group.
    - Path Parameters:
        - `group_id`: The ID of the group to remove the user from.
        - `member_id`: The ID of the user to remove.
    - Response: 
        - Status: `204 No Content`
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid user ID"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "User not found in group"
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```


---

- `/my-groups`:
    - Method: `GET`
    - Description: Retrieves a list of all groups the current user is a member of.
    - Query Parameters:
        - `page_size`: The page number for pagination.
        - `page_number`: The number of groups per page.  (1 - 200)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data" {
                "page_size": 10,
                "page_number": 1,
                "total_pages": 100,
                "prev": 1 or null,
                "next": 3 or null,
                "data": [
                    {
                        "group_id": "string",
                        "name": "string",
                        ...,
                        "is_admin": true or false // (computed by the server)
                    },
                    ...
                ]`
            }
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid page size or number"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Group not found"
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```    

---

## Group Transactions:
*Prefix*: `/groups/{group_id}/transactions`
- *Common response*
    - `200 OK`
    - `201 Created`
    - `204 No Content`
    - `400 Bad Request`
    - `401 Unauthorized`
    - `403 Forbidden`
    - `404 Not Found`
    - `409 Conflict`: User not in group

- `/`:
    - Method: `GET`
    - Description: Retrieves a list of all transactions.
    - Query Parameters:
        - `page_size`: The page number for pagination.
        - `page_number`: The number of transactions per page.  (1 - 200)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data" {
                "page_size": 10,
                "page_number": 1,
                "total_pages": 100,
                "prev": 1 or null,
                "next": 3 or null,
                "data": [
                    {
                        "transaction_id": "string",
                        ...
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid page size or number"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Transaction not found"
        }
        ```
- `/`:
    - Method: `POST`
    - Description: Creates a new transaction.
    - Request Body: 
    ```json
    {
        "amount": 100,
        "description": "string",
        ...
    }
    ```
    - Response: 
        - Status: `201 Created`
        - Body: 
        ```json
        {
            "data": {
                "transaction_id": "string",
                ...
                "amount": 100 or null,
                "participations": [
                    {
                        "user_id": "string",
                        "amount": 100 or null,
                        "description": "string"
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid transaction amount"
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
- `/{transaction_id}`
    - Method: `GET`
    - Description: Retrieves a specific transaction's information.
    - Path Parameters:
        - `transaction_id`: The ID of the transaction to retrieve.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "transaction_id": "string",
                ...
                "amount": 100 or null,
                "participations": [
                    {
                        "user_id": "string",
                        "amount": 100 or null,
                        "description": "string"
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Transaction not found"
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
- `/{transaction_id}`
    - Method: `PUT`
    - Description: Updates a specific transaction's information.
    - Path Parameters:
        - `transaction_id`: The ID of the transaction to update.
    - Request Body:
    ```json
    {
        "amount": 100,
        "description": "string",
        "amount": 100 or null,
        "participations": [
            {
                "user_id": "string",
                "amount": 100 or null,
                "description": "string"
            },
            ...
        ],
        ...
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "transaction_id": "string",
                ...
                "amount": 100 or null,
                "participations": [
                    {
                        "user_id": "string",
                        "amount": 100 or null,
                        "description": "string"
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
    ...
- `/{transaction_id}/description`:
    - Method: `PATCH`
    - Description: Partially updates a specific transaction's description.
    - Path Parameters:
        - `transaction_id`: The ID of the transaction to update.
    - Query Parameters:
        - `description`: The new description of the transaction.
    - Response:
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "transaction_id": "string",
                ...
                "amount": 100 or null,
                "participations": [
                    {
                        "user_id": "string",
                        "amount": 100 or null,
                        "description": "string"
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{transaction_id}`:
    - Method: `DELETE`
    - Description: Deletes a specific transaction.
    - Path Parameters:
        - `transaction_id`: The ID of the transaction to delete.
    - Response: 
        - Status: `204 No Content`
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Transaction not found"
        }
        ```
- `/{transaction_id}/participations/`:
    - Method: `POST`
    - Description: Adds some users to a specific transaction.
    - Path Parameters:
        - `transaction_id`: The ID of the transaction to add the users to.
    - Request Body:
    ```json
    {
        "description": "string",
        "amount": 100 or null,
        "participations": [
            {
                "user_id": "string",
                "amount": 100 or null,
                "description": "string"
            },
            ...
        ],
    }
    ```
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "transaction_id": "string",
                ...
                "amount": 100 or null,
                "participations": [
                    {
                        "user_id": "string",
                        "amount": 100 or null,
                        "description": "string"
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/{transaction_id}/participations/{participation_id}`:
    - Method: `DELETE`
    - Description: Removes a specific user from a transaction.
    - Path Parameters:
        - `transaction_id`: The ID of the transaction to remove the user from.
        - `participation_id`: The ID of the user to remove.
    - Response: 
        - Status: `204 No Content`
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid user ID"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "User not found in transaction"
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
    - Error Response:
        - Status: `409 Conflict`
        - Body: 
        ```json
        {
            "error": "User not in group"
        }
        ```
---

## Debt Calculation:
*Prefix*: `/groups/{group_id}`
- *Common response*
    - `200 OK`
    - `201 Created`
    - `204 No Content`
    - `400 Bad Request`
    - `401 Unauthorized`
    - `403 Forbidden`
    - `404 Not Found`
    - `409 Conflict`: User not in group

- `/debt-calculation`:
    - Method: `GET`
    - Description: Retrieves the debt calculation for a specific group.
    - Path Parameters:
        - `group_id`: The ID of the group to retrieve the debt calculation for.
        - `page_size`: The page number for pagination.
        - `page_number`: The number of transactions per page.  (1 - 200)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data" {
                "page_size": 10,
                "page_number": 1,
                "total_pages": 100,
                "prev": 1 or null,
                "next": 3 or null,
                "data": [
                    {
                        "user_id": "string",
                        "debt": 100 or null,
                        ...
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid page size or number"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Group not found"
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```
- `/pay-debt`:
    - Method: `POST`
    - Description: Settles the debts for a specific group.
    - Request Body: 
    ```json
    {
        "payer_id": "string",
        "payee_id": "string",
        "amount": 100,
        "description": "string",
        ...
    }
    ```
    - Response: 
        - Status: `204 No Content`
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid transaction amount"
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "User not found"
        }
        ```
    - Error Response:
        - Status: `409 Conflict`
        - Body: 
        ```json
        {
            "error": "User not in group"
        }
        ```
- `/balance`:
    - Method: `GET`
    - Description: Retrieves the balance for a specific group.
    - Path Parameters:
        - `group_id`: The ID of the group to retrieve the balance for.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "user_id": "string",
                "balance": 100,
                ...
            }
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Group not found"
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid group ID"
        }
        ```
- `/my-balance`:
    - Method: `GET`
    - Description: Retrieves the balance for the current user.
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "user_id": "string",
                "balance": 100,
                ...
            }
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid user ID"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "User not found"
        }
        ```

---

## My Balances:
*Prefix*: `/my-balances`
- *Common response*
    - `200 OK`
    - `201 Created`
    - `204 No Content`
    - `400 Bad Request`
    - `401 Unauthorized`
    - `403 Forbidden`
    - `404 Not Found`
    - `409 Conflict`: User not in group
- `/`:
    - Method: `GET`
    - Description: Retrieves a list of all balances for the current user in all groups.
    - Query Parameters:
        - `page_size`: The page number for pagination.
        - `page_number`: The number of transactions per page.  (1 - 200)
    - Response: 
        - Status: `200 OK`
        - Body: 
        ```json
        {
            "data": {
                "page_size": 10,
                "page_number": 1,
                "total_pages": 100,
                "prev": 1 or null,
                "next": 3 or null,
                "data": [
                    {
                        "group_id": "string",
                        "balance": 100,
                        "debts": [
                            {
                                "user_id": "string",
                                "amount": 100,
                                ...
                            },
                            ...
                        ],
                        "credits": [
                            {
                                "user_id": "string",
                                "amount": 100,
                                ...
                            },
                            ...
                        ]
                        ...
                    },
                    ...
                ]
            }
        }
        ```
    - Error Response:
        - Status: `401 Unauthorized`
        - Body: 
        ```json
        {
            "error": "Unauthorized"
        }
        ```
    - Error Response:
        - Status: `400 Bad Request`
        - Body: 
        ```json
        {
            "error": "Invalid page size or number"
        }
        ```
    - Error Response:
        - Status: `404 Not Found`
        - Body: 
        ```json
        {
            "error": "Group not found"
        }
        ```
    - Error Response:
        - Status: `403 Forbidden`
        - Body: 
        ```json
        {
            "error": "Permission denied"
        }
        ```