import asyncio

def function_that_returns_a_future_object():
    print("future object function~!")
    loop = asyncio.get_event_loop()
    loop.call_soon(lambda : print("hello world"))
    future = loop.create_future()
    future.set_result("success")
    return future

async def main():
    result = await function_that_returns_a_future_object()
    print(result)
    print("AAAAAAAA")
    print("BBBBBBBB")
    print("CCCCCCCC")

# loop = asyncio.get_event_loop()
# loop.run_until_complete(main())
# loop.close()
asyncio.run(main())