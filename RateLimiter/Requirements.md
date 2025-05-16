## Rate Limiter

### Requirements:
1. System should be able to set rules for throttling different APIs with respect to either userId or IPaddress.
2. System should support multiple rate limiting algorithms.
3. Rule should be based on Noof Requests allowed per minute.
4. Rules can be changed dynamically and next requests should be handled according to new ruled.

### Core Components
1. ThrottleConfig
2. RequestStore
3. RateLimiterInterface : Rate Limiter Factory

### Design Patterns
1. Strategy Pattern : To apply different rate limit algorithms
2. Factory Pattern : To get rateLimiter Algo