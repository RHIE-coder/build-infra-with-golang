# DynamoDB

DynamoDB는 스키마 없는 NoSQL 데이터베이스입니다. 이는 테이블을 생성할 때 기본 키 속성 외에는 다른 속성이나 데이터 형식을 정의할 필요가 없다는 의미입니다. 

## [ 구성요소 ]
 - 테이블
 - 속성
 - 항목

## [ 속성: 데이터 유형 ]

 - 스칼라 형식: 스칼라 형식은 하나의 값만 표현할 수 있습니다. 스칼라 형식은 숫자, 문자열, 이진수, 부울 및 Null입니다.
 - 문서 형식: 문서 형식은 중첩된 속성이 있는 복잡한 구조를 표현할 수 있습니다. 이러한 형식은 JSON 문서에서 찾을 수 있습니다. 문서 형식은 목록 및 맵입니다.
   - 목록: `["A", "B", 3.14]`
   - 멥: `{name: "rhie", age: 20}`
 - 집합 형식: 집합 형식은 여러 스칼라 값을 표현할 수 있습니다. 집합 형식은 문자열 집합, 숫자 집합 및 이진수 집합입니다.

## [ API ]

 - PutItem:  항목을 생성합니다.
 - GetItem:  항목을 읽습니다.
 - UpdateItem:  항목을 업데이트합니다.
 - DeleteItem:  항목을 삭제합니다.

네 개의 기본 CRUD 작업에 추가하여 DynamoDB는 다음과 같은 작업도 제공합니다.

 - BatchGetItem : 하나 이상의 테이블에서 최대 100개의 항목을 읽습니다.
 - BatchWriteItem : 하나 이상의 테이블에서 최대 25개의 항목을 생성하거나 삭제합니다.

## [ API - 항목 읽기 ]

### - 기본

기본적으로 GetItem 요청은 최종적으로 일관된 읽기를 수행

>```
>aws dynamodb get-item \
>    --table-name ProductCatalog \
>    --key '{"Id":{"N":"1"}}'
>```

###  - 옵션

 - ConsistentRead 파라미터를 사용하여 강력한 일관된 읽기를 요청
 - projection-expression을 사용하여 속성의 일부만 반환
 - GetItem에서 사용된 읽기 용량 단위 수를 반환하려면 ReturnConsumedCapacity 파라미터를 TOTAL로 설정

>```
>aws dynamodb get-item \
>    --table-name ProductCatalog \
>    --key '{"Id":{"N":"1"}}' \
>    --consistent-read \
>    --projection-expression "Description, Price, RelatedItems" \
>    --return-consumed-capacity TOTAL
>```

## [ API - 항목 쓰기 ]

사용된 쓰기 용량 단위 수를 반환하려면 ReturnConsumedCapacity 파라미터 설정

 - TOTAL : 사용된 총 쓰기 용량 단위 수를 반환합니다.
 - INDEXES : 사용된 총 쓰기 용량 단위 수와 작업의 영향을 받은 테이블 및 보조 인덱스의 소계를 반환합니다.
 - NONE : 쓰기 용량 세부 정보를 반환하지 않습니다. (이 값이 기본값입니다.)

### - PutItem

새 항목을 만듬. 동일 키의 항목이 테이블이 이미 존재하는 경우 해당 항목이 새 항목으로 대체.

#### 기본

>```
>aws dynamodb put-item \
>    --table-name Thread \
>    --item file://item.json
>```

#### ReturnValues

 - ReturnValues: ALL_OLD

기존 항목을 덮어쓰면 ALL_OLD는 덮어쓰기 전에 나타난 전체 항목을 반환합니다.

존재하지 않는 항목을 쓰면 ALL_OLD는 효과를 나타내지 않습니다.

### - UpdateItem

지정된 키가 있는 항목이 존재하지 않으면 UpdateItem은 새 항목을 생성, 그렇지 않으면 기존 항목의 속성을 수정

#### 기본

>```
>aws dynamodb update-item \
>    --table-name Thread \
>    --key file://key.json \
>    --update-expression "SET Answered = :zero, Replies = :zero, LastPostedBy = :lastpostedby" \
>    --expression-attribute-values file://expression-attribute-values.json \
>    --return-values ALL_NEW
>```

#### ReturnValues

UpdateItem의 가장 일반적인 용도는 기존 항목을 업데이트하는 것입니다. 하지만 실제로 UpdateItem은 항목이 아직 없는 경우 항목을 자동으로 생성하는 upsert를 수행합니다.

 - ReturnValues: ALL_OLD

기존 항목을 업데이트하면 ALL_OLD는 업데이트 전에 나타난 전체 항목을 반환합니다.

존재하지 않는 항목을 업데이트하면(upsert) ALL_OLD는 효과를 나타내지 않습니다.

 - ReturnValues: ALL_NEW

기존 항목을 업데이트하면 ALL_NEW는 업데이트 후에 나타난 전체 항목을 반환합니다.

존재하지 않는 항목을 업데이트하면(upsert) ALL_NEW는 전체 항목을 반환합니다.

 - ReturnValues: UPDATED_OLD

기존 항목을 업데이트하면 UPDATED_OLD는 업데이트 전에 나타난 업데이트된 속성만 반환합니다.

존재하지 않는 항목을 업데이트하면(upsert) UPDATED_OLD는 효과를 나타내지 않습니다.

 - ReturnValues: UPDATED_NEW

기존 항목을 업데이트하면 UPDATED_NEW는 업데이트 후에 나타난 영향을 받은 속성만 반환합니다.

존재하지 않는 항목을 업데이트하면(upsert) UPDATED_NEW는 업데이트 후에 나타나는 업데이트된 속성만 반환합니다.

### - DeleteItem

#### 기본

>```
>aws dynamodb delete-item \
>    --table-name Thread \
>    --key file://key.json
>```

#### ReturnValues

- ReturnValues: ALL_OLD

기존 항목을 삭제하면 ALL_OLD는 삭제하기 전에 나타난 전체 항목을 반환합니다.

존재하지 않는 항목을 삭제하면 ALL_OLD는 데이터를 반환하지 않습니다.

