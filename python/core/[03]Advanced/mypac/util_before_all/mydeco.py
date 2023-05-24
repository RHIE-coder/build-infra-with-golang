class Deco:
    def __init__(self, func):
        self.func = func
        print("DECO!!!")
        
    def __call__(self, a, b):
        print(f'매개변수 정보 {a}, {b}')
        return self.func(a,b)

class Paradeco:
    def __init__(self, nums):
        self.nums = nums
        
    def __call__(self, func):
        def wrapper(*, do):
            doWhat = None
            param = dict(do = do)
            if do == "add":
                doWhat = lambda x : sum(x)
            elif do == "max":
                doWhat = lambda x : max(x)
            func(**param)
            return doWhat(self.nums) if doWhat != None else "Invalid"
        return wrapper