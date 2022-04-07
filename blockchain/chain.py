class Block:
    def __init__(self, data, prev=None):
        if prev:
            prev.next = self
        self.prev = prev
        self.next = None
        self.data = data


first_block = Block(0)
chain = first_block
for i in range(1, 10):
    first_block = Block(i, first_block)

while chain.next != None:
    print(chain.data)
    chain = chain.next
