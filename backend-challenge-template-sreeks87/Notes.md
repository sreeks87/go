Manipulation of state
If a user manipulates tDB connection that could store the previous state - position of fruit and snake, score etc
This can be further used toverify the request that is recreived by the system.
Currenlty the POST request is not using the already created Games, therefore most of the POST would 
result in a NOTFOUND in DB error, so I have disabled the check in this final commit.

API versioning
the current service cant be versioned, mainly due to the test binary expecting the endpoints to be ina specific format.

I am purposefully using HTTP codes in the service layer, this is not a standard practice
the service should not know what kind of request it is serving at the moment.
The handler layer should be unknow to service layer. But this is being done for simplicity at the moment.


== PLEASE INCLUDE THIS WITH YOUR SUBMISSION. ==
Well done! Your code is: hacker-c653vs1vm9g1u353j5n0-man
== PLEASE INCLUDE THIS WITH YOUR SUBMISSION. ==