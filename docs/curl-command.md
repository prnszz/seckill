1. 商品管理API：
```bash
# 1. 创建商品
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "description": "最新款苹果手机",
    "price": 5999.00,
    "stock": 100,
    "status": 1
  }'

# 2. 获取商品详情
curl http://localhost:8080/api/v1/products/1

# 3. 更新商品
curl -X PUT http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "iPhone 15",
    "description": "最新款苹果手机 - 更新描述",
    "price": 5899.00,
    "stock": 200,
    "status": 1
  }'

# 4. 删除商品
curl -X DELETE http://localhost:8080/api/v1/products/1

# 5. 获取商品列表
curl "http://localhost:8080/api/v1/products?page=1&page_size=10"
```

2. 秒杀活动管理API：
```bash
# 1. 创建秒杀活动
curl -X POST http://localhost:8080/api/v1/seckill/activities \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "seckill_price": 4999.00,
    "seckill_stock": 50,
    "start_time": "2024-12-16T10:00:00Z",
    "end_time": "2024-12-16T12:00:00Z",
    "status": 0
  }'

# 2. 获取秒杀活动详情
curl http://localhost:8080/api/v1/seckill/activities/1

# 3. 更新秒杀活动
curl -X PUT http://localhost:8080/api/v1/seckill/activities \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "seckill_price": 4899.00,
    "seckill_stock": 100,
    "start_time": "2024-12-16T10:00:00Z",
    "end_time": "2024-12-16T12:00:00Z"
  }'

# 4. 获取活动列表（所有状态）
curl "http://localhost:8080/api/v1/seckill/activities?page=1&page_size=10"

# 5. 获取指定状态的活动列表
# status: 0-未开始，1-进行中，2-已结束
curl "http://localhost:8080/api/v1/seckill/activities?page=1&page_size=10&status=1"
```

3. 秒杀核心API：
```bash
# 1. 预热商品库存
curl -X POST "http://localhost:8080/api/v1/seckill/preload/1?stock=100"

# 2. 执行秒杀
curl -X POST http://localhost:8080/api/v1/seckill/do/1
```

4. 数据库连接测试：
```bash
# 测试数据库连接
curl http://localhost:8080/test/db
```

使用这些命令时需要注意：
1. 请确保先创建商品，再创建秒杀活动
2. 创建秒杀活动时，product_id 必须是已存在的商品ID
3. 执行秒杀前，需要先预热库存
4. 时间字段使用 UTC 格式，注意调整为你需要的时间
5. 所有涉及 ID 的请求（如 /products/1 或 /activities/1），请使用实际存在的 ID

你可以按以下顺序测试完整流程：
1. 测试数据库连接
2. 创建商品
3. 创建秒杀活动
4. 预热库存
5. 执行秒杀

