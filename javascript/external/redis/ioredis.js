const Redis = require("ioredis");

(async()=>{
    const redis = new Redis({
        host: "192.168.154.1",
        port: 6379,
    });
    
    
    await redis.set("greeting", "Hello World");
    console.log(await redis.get("greeting"));
    await redis.expire("greeting", 10);

    redis.disconnect();
})() 