class Solution:

    def encode(self, strs: List[str]) -> str:
        key = "split_key"
        decoded_list = []
        for s in strs:
            decoded_list.append(s + "split_key")
        return "".join(decoded_list)

    def decode(self, s: str) -> List[str]:
        return s.split("split_key")[:-1]
