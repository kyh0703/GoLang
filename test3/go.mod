module test3

go 1.18

replace (
	event => ./event
	kafkaevent => ./event/kafkaevent
)
