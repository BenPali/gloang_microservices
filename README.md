# golang_microservices

Project:

The main goal of the project is to build the back-end architecture of a digital marketplace.
The project will be written in Golang and must use microservices.

How you separate your microservices is up to you, and you will be evaluated on this.

Your back-end must handle accounts, ads, transactions.

For all the informations bellow you are free to use the architecture that you think will be the best.
For example, if I'm describing an account with it's mandatory informations and mandatory actions, that doesn't mean that you are forced to use a single structure to handle the informations or a single microservices to handle the actions.

+----------+
| Accounts |
+----------+

An account should have at least the following informations (you can add some if needed):
    - email (Must be unique)
    - login (Must be unique)
    - password
    - balance

Your back-end must implement the following actions related to an account (you can add some if needed):
    - Create an account
    - Login
    - Update informations of its own account
    - Delete its own account
    - Fully read its own account
    - Partially read any user account
    - Add funds to it's own balance

+----_+
| Ads |
+-----+

An ad should be linked to an account

An ad should have at least the following informations (you can add some if needed):
    - title
    - description
    - price
    - picture

Your back-end must implement the following actions related to an ad (you can add some if needed):
    - Create an ad
    - Update one of its own ad
    - Delete one of its own ad
    - Read any ad
    - Get a list of ads searching by keywords
    - Get a list of all the ads of a user

+--------------+
| Transactions |
+--------------+

A transaction should be linked to two accounts and an ad

A transaction should have at least the following informations (you can add some if needed):
    - messages
    - bid prices
    - status

Your back-end must implement the following actions related to a transaction (you can add some if needed):
    - Make an offer on an ad
    - Accept an offer on its own ad
    - Refuse an offer on its own ad
    - List all its own transaction

+------+
| Hint |
+------+

- Some API requests must be called with an authenticated users, others don't (Be pragmatic).
- You might need to authorize user's on some API requests (Think about access token).
- Some data might be sensible and should not be shared to everyone (One's again, be pragmatic).
- Obviously making an offer and accepting a transaction has several impact on data.
- Think of a way to avoid double spending problem.
- Searching an object in a database based on a subquery string isn't a good idea.
- Micro services needs to communicate between each others but not with the same reliability.

+-------+
| Bonus |
+-------+

Add an admin role. Admin should be able to:
    - Update / delete any account
    - Fully read any account
    - Add / retrieve funds to any account
    - Delete any ads
    - Cancel any transaction (think about the consequence of a transaction cancel)
