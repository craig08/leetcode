class Solution:
    def equationsPossible(self, equations: List[str]) -> bool:
        j = 0
        for i, eq in enumerate(equations):
            if eq[1] == '=':
                equations[i], equations[j] = equations[j], equations[i]
                j += 1
        self.groups, self.ranks = dict(), dict()
        for eq in equations:
            if eq[0] not in self.groups:
                self.groups[eq[0]] = eq[0]
                self.ranks[eq[0]] = 1
            if eq[-1] not in self.groups:
                self.groups[eq[-1]] = eq[-1]
                self.ranks[eq[-1]] = 1
            g1, g2 = self.find(eq[0]), self.find(eq[-1])
            if g1 == g2 and eq[1] == '!':
                return False
            elif g1 != g2 and eq[1] == '=':
                self.union(g1, g2)
        return True

    def find(self, var) -> int:
        if self.groups[var] == var:
            return var
        self.groups[var] = self.find(self.groups[var])
        return self.groups[var]

    def union(self, g1, g2):
        if self.ranks[g1] < self.ranks[g2]:
            self.groups[g1] = g2
            self.ranks[g2] += self.ranks[g1]
            del self.ranks[g1]
        else:
            self.groups[g2] = g1
            self.ranks[g1] += self.ranks[g2]
            del self.ranks[g2]
