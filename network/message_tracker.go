package network

import (
	"errors"

	"github.com/emirpasic/gods/maps/treemap"
)

// MessageTracker tracks a configurable fixed amount of messages.
// Messages are stored first-in-first-out.  Duplicate messages should not be stored in the queue.
type MessageTracker interface {
	// Add will add a message to the tracker, deleting the oldest message if necessary
	Add(message *Message) (err error)
	// Delete will delete message from tracker
	Delete(id string) (err error)
	// Get returns a message for a given ID.  Message is retained in tracker
	Message(id string) (message *Message, err error)
	// Messages returns messages in FIFO order
	Messages() (messages []*Message)
}

// ErrMessageNotFound is an error returned by MessageTracker when a message with specified id is not found
var ErrMessageNotFound = errors.New("message not found")
var ErrReachMaxLength = errors.New("container reach max length")

type MessageTrackerImp struct {
	// MessageTracker
	max_len  int
	messages *treemap.Map
}

func (mt MessageTrackerImp) Add(message *Message) (err error) {
	if _, found := mt.messages.Get(message.ID); found {
		return nil
	}

	length := len(mt.messages.Keys())
	if length == mt.max_len {
		min, _ := mt.messages.Min()
		mt.messages.Remove(min)
	}

	mt.messages.Put(message.ID, message)
	return nil
}

func (mt MessageTrackerImp) Delete(id string) (err error) {
	if _, found := mt.messages.Get(id); !found {
		return ErrMessageNotFound
	} else {
		mt.messages.Remove(id)
		return nil

	}
}

func (mt MessageTrackerImp) Message(id string) (message *Message, err error) {
	if message, found := mt.messages.Get(id); !found {
		return nil, ErrMessageNotFound
	} else {
		return message.(*Message), nil
	}
}

func (mt MessageTrackerImp) Messages() (messages []*Message) {
	interfaces := mt.messages.Values()
	for _, v := range interfaces {
		messages = append(messages, v.(*Message))
	}
	return messages
}

func NewMessageTracker(length int) MessageTracker {
	// TODO: Implement this constructor with your implementation of the MessageTracker interface
	return MessageTrackerImp{
		max_len:  length,
		messages: treemap.NewWithStringComparator(),
	}

}
