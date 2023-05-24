from functools import reduce
from collections import deque

class HigherOrder:
    def __init__(self, *args):
        self.args = args
        self.isVerboseMode = False

    def verbose(self, isVerboseMode):
        self.isVerboseMode = isVerboseMode
    
    def changeValues(self, *args):
        self.args = args

    def functionsChain(self, *, mapping = None, filtering = None, reducing = None):
        self.isVerboseMode if print("Parsing Target : ", self.args) else None
        queue = deque([mapping, filtering, reducing])
        self.isVerboseMode if print("Parsing Steps : ", queue) else None
        source = list(self.args)
        x = queue.popleft()
        self.isVerboseMode if print("Mapping Steps : ", x) else None
        if x != None:
            source = list(map(x,source))
            self.isVerboseMode if print("Mapping After: ", source) else None
        x = queue.popleft()
        self.isVerboseMode if print("Filtering Steps : ", x) else None
        if x != None:
            source = filter(x, source)
            self.isVerboseMode if print("Filtering After: ", source) else None
        x = queue.popleft()         
        self.isVerboseMode if print("Reducing Steps : ", x) else None
        if x != None:
            source = reduce(x, source)
            self.isVerboseMode if print("Reducing After: ", source) else None
                
        self.isVerboseMode if print("Result : ", source) else None
        
        return source