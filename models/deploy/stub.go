package deploy

import (
	"pilot/models/deploy/board"
)

type stubStore struct{

}

func NewBoardStore() (board.BoardStore, error) {
	is := &stubStore{
	}

	return is, nil
}

func (is *stubStore) Walk(f board.BoardWalkFunc) error {
	return nil
}

func (is *stubStore) Store(name string, board *board.Board) error{
	return nil
}

func (is *stubStore) Get(name string) (*board.Board, error) {
	return nil, nil
}

func (is *stubStore) Delete(name string) error {
	return nil
}

func (is *stubStore) Update(name string, board *board.Board) error{
	return nil
}
