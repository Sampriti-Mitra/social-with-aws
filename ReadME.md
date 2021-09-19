# Social App

## I. Set up
### Pre Requisites
1. Working AWS account
2. Mysql local db

### How to set up
1. Run migrations from the internal/migrations folder
2. Add the Key and Secret of AWS account in the config/token.go or export as env variable
3. Add the db credentials in config/token.go or export as env variable
4. Add the bucket name in config/token.go

## II. Run the postman collection
Find the postman collection in the root of the directory social.postman_collection.json
and import it in postman

## III. What to expect from this project
1. Create an account
2. Create a post from an account, send header as the account-id
3. Write a comment from an account on a post
4. Fetch all posts from all accounts



