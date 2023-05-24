import asyncio

async def saying(interval, what, how_many):
    while how_many > 0:
        await asyncio.sleep(interval)
        print(what)
        how_many -= 1

async def main():
    coro1 =  saying(1, 'hello', 5)
    coro2 =  saying(2, 'world', 3)

    await coro1
    await coro2

asyncio.run(main())