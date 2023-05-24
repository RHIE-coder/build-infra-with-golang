import asyncio
import time

async def say_after(delay, what):
    await asyncio.sleep(delay)
    print(what)

async def main():
    coro1 =  say_after(3, 'hello')
    coro2 =  say_after(4, 'world')

    task1 = asyncio.create_task(coro1)
    task2 = asyncio.create_task(coro2)

    print(f"started at {time.strftime('%X')}")

    await task1
    await task2

    print(f"finished at {time.strftime('%X')}")

asyncio.run(main())