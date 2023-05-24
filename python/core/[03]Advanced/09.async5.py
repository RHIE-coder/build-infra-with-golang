import asyncio

async def saying(interval, what, how_many):
    while how_many > 0:
        await asyncio.sleep(interval)
        print(what)
        how_many -= 1

async def main():
    coro1 =  saying(1, 'hello', 5)
    coro2 =  saying(2, 'world', 3)
    coro3 =  saying(3, 'python', 2)

    task1 = asyncio.create_task(coro1)
    task2 = asyncio.create_task(coro2)
    task3 = asyncio.create_task(coro3)

    await task1
    await task2
    await task3

asyncio.run(main())