class Solution:
    def validIPAddress(self, IP: str) -> str:
        if '.' in IP:
            groups = IP.split('.')
            if len(groups) != 4:
                return 'Neither'
            for num in groups:
                try:
                    if not 0 < len(num) < 4:
                        return 'Neither'
                    if (int(num) == 0 or num[0] == '0') and len(num) > 1:
                        return 'Neither'
                    if not 0 <= int(num) < 256:
                        return 'Neither'
                except:
                    return 'Neither'
            return 'IPv4'
        elif ':' in IP:
            groups = IP.split(':')
            if len(groups) != 8:
                return 'Neither'
            for group in groups:
                if len(group) == 0 or len(group) > 4:
                    return 'Neither'
                if any(g for g in group if g.lower() not in '0123456789abcdef'):
                    return 'Neither'
            return 'IPv6'
        return 'Neither'
