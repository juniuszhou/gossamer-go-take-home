# Design for MessageTracker

The homework requires a structure to store network messages, you can add/delete/get/get all. With the limited size.

From the test cases, we can see some details
1. if add a message with same id, just keep the old one, no replacement
2. return err when remove a non-existed message with id
3. return message or nil when get message by id
4. return all messages in order when get all messages
5. if add a message but the container is full, we should remove the oldest one and add new one
6. sort messages according to id

## Solution
To get the message quickly and sort message efficiently, we can choose a sorted map or heap as underlying container.
And choose the treemap from github.com/emirpasic/gods.

## Test 
unit tests cover all source code. 
