import asyncio

async def saying(interval, what, how_many):
    while how_many > 0:
        await asyncio.sleep(interval)
        print(what)
        how_many -= 1
    return what

async def main():
    coro1 =  saying(1, 'hello', 5)
    coro2 =  saying(2, 'world', 3)
    coro3 =  saying(3, 'python', 2)

    coro_result_list = await asyncio.gather(
        coro1, coro2, coro3
    )
    print(coro_result_list)

asyncio.run(main())