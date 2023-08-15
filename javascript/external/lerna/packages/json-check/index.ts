import Ajv from 'ajv';
const ajv = new Ajv();

export function isValid(data:any, schema:any): boolean {
    const valid:boolean = ajv.validate(schema, data)
    return valid
}