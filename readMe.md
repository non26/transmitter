this project is trying to receive http request from trading view hook via API gateway and the push the reqeust to SQS and SQS will trigger the bot lambda function.

---> Problem To Solve
this project is aim to solve the problem when receiving the request from trading view hook, the request will have the time out for 3 second. When sending the request to Bot sometimes the Bot takes processing more than 3 second, so the request is fail and bot make no position.

----> Solving
push the request to sqs and so that the bot can processing the request longer that 3 second.