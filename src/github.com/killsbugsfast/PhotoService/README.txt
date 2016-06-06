Things We Didn’t Do
-------------------
While we are off to a great start, there is a lot left to do. Things we haven’t addressed are:

Version Control – What if we need to modify the API and that results in a breaking change? Might we add /v1/prefix to all our routes to start with?
Authentication – Unless this is a free/public API, we probably need some authentication. I suggest learning about JSON web tokens
eTags – If you are building something that needs to scale, you will likely need to implement eTags


Command to add a new json object into the database
--------------------------------------------------
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

git commands
------------
git status     
git add -A     // add all files
git commit -m "My check-in message goes here"
git push origin master // push the change to master