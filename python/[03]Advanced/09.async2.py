import asyncio
import time

async def say_after(delay, what):
    await asyncio.sleep(delay)
    print(what)

async def main():
    # https://docs.python.org/3.8/library/time.html#time.strftime
    print(f"started at {time.strftime('%X')}") # Locale’s appropriate time representation.

    await say_after(3, 'hello')
    await say_after(4, 'world')

    print(f"finished at {time.strftime('%X')}")

asyncio.run(main())