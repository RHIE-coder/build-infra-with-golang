(async():Promise<unknown>=>{

    type a = Awaited<Promise<string>>

    // all properties of Type set to optional.
    interface Todo {
        title: string;
        description: string;
    }

    type b = Partial<Todo>

    // all properties of Type set to required.
    interface Props {
        a?: number;
        b?: string;
    }

    type c = Required<Props>

    // Readonly<Type>
    // Record<Keys, Type> key-value pair
    // Pick<Type, Keys> Type으로부터 K 프로퍼티만 추출
    // Omit<T,K> T에서 모든 프로퍼티를 선택한 다음 K를 제거한 타입을 구성
    // Exclude<UnionType, ExcludedMembers> T에서 U에 할당할 수 있는 모든 속성을 제외한 타입을 구성
    // Extract<Type, Union> 
    // NonNullable<Type>
    // Parameters<Type> 함수 타입 T의 매개변수 타입들의 튜플 타입을 구성
    function func(arg: {a: number, b: string}): void{}
    type f0 = typeof func
    type f1 = Parameters<typeof func>

    // ConstructorParameters<Type>
    // ReturnType<Type>
    // InstanceType<Type>
    // ThisParameterType<Type> this 매개변수가 없을 경우 unknown을 추출
    // OmitThisParameter<Type> 함수 Type에서 'this' 매개변수를 제거 --strictFunctionTypes가 활성화되었을 때만 올바르게 동작
    // ThisType<Type> noImplicitThis 의 플래그가이 유틸리티를 사용하도록 설정


    /*  
        Uppercase<StringType>
        Lowercase<StringType>
        Capitalize<StringType>
        Uncapitalize<StringType>
    */
    return
})()