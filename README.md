# “Word of Wisdom quotes” tcp server

Task: Design and implement
• TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), 
the challenge-response protocol should be used.
• The choice of the POW algorithm should be explained.
• After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other 
collection of the quotes.
• Docker file should be provided both for the server and for the client that solves the POW Challenge.

Solution: 
* In case of content delivering POW don't protect from DDOS attacks, because we make small active work for each request.
  (Checking of POW result harder than delivering of random static content even without of static cashing like cloudfront etc.)
* Chosen POW algorithm is "Hashcash" (https://en.wikipedia.org/wiki/Hashcash) because it is simple and easy to implement.
* Implementation simplified. We expect, that Js client will calculate correct integer nonce, that counted via a hash of last Job
  and timestamp of it, and provides given hash. 
  No zeros or unknown correct answers, and etc. Just "try to gas what the server generate".
  (we have no consensus, so we don't need to make it more complex)
* if correct answers came too fast, server will increase difficulty of the task.
* From the UI side expected a hiding the POW logic.
* The project don't have a CI/CD, so docker container will be built manually and compile the code by them self.

Hashing scheme for resolving the POW job:
Sha1("last hash" + "number" + "created")
Where: 
* "last hash" - hash of the previous job
* "number" - integer number, that will be found by the client
* "created" - timestamp of the job creation

Difficulty it's a max number that was used for generating the "number" field.