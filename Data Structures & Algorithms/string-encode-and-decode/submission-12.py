class Solution:

    def encode(self, strs: List[str]) -> str:
        for i, s in enumerate(strs):
            strs[i] = f"{len(s)}#{s}"
        return "".join(strs)

    def decode(self, s: str) -> List[str]:
        result = []
        l = ""
        i = 0
        while i < len(s):
            if s[i] == "#":
                start = i + 1
                end = start + int(l)
                result.append(s[start:end])
                i = end
                l = ""
            else:
                l += s[i]
                i += 1
        return result
