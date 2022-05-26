module test3

go 1.17

replace (
	event => ./event
	kafkaevent => ./event/kafkaevent
)

require event v0.0.0-00010101000000-000000000000

require kafkaevent v0.0.0-00010101000000-000000000000 // indirect
