class Solution:
    def updateBoard(self, board: List[List[str]], click: List[int]) -> List[List[str]]:
        # if it is a mine
        if board[click[0]][click[1]] == 'M':
            board[click[0]][click[1]] = 'X'
            return board
        self.reveal(board, click)
        return board

    def countMines(self, board: List[List[str]], click: List[int]) -> int:
        diff = [(x, y) for x in range(-1, 2) for y in range(-1, 2)]
        count = 0
        for d in diff:
            if d[0] == d[1] == 0:
                continue
            point = [click[0]+d[0], click[1]+d[1]]
            if point[0] < 0 or point[1] < 0 or point[0] >= len(board) or point[1] >= len(board[0]):
                continue
            if board[point[0]][point[1]] == 'M':
                count += 1
        return count

    def reveal(self, board: List[List[str]], click: List[int]):
        mines = self.countMines(board, click)
        if mines > 0:
            board[click[0]][click[1]] = '{}'.format(mines)
            return

        board[click[0]][click[1]] = 'B'
        diff = [(x, y) for x in range(-1, 2) for y in range(-1, 2)]
        for d in diff:
            point = [click[0]+d[0], click[1]+d[1]]
            if point[0] < 0 or point[1] < 0 or point[0] >= len(board) or point[1] >= len(board[0]) or board[point[0]][point[1]] != 'E':
                continue
            self.reveal(board, point)
