// @ts-nocheck
/* eslint-disable */
/**
 * 
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SchemasUpdateMessageStatusPayload
 */
export interface SchemasUpdateMessageStatusPayload {
    /**
     * 
     * @type {boolean}
     * @memberof SchemasUpdateMessageStatusPayload
     */
    isRead: boolean;
}

/**
 * Check if a given object implements the SchemasUpdateMessageStatusPayload interface.
 */
export function instanceOfSchemasUpdateMessageStatusPayload(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "isRead" in value;

    return isInstance;
}

export function SchemasUpdateMessageStatusPayloadFromJSON(json: any): SchemasUpdateMessageStatusPayload {
    return SchemasUpdateMessageStatusPayloadFromJSONTyped(json, false);
}

export function SchemasUpdateMessageStatusPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasUpdateMessageStatusPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'isRead': json['isRead'],
    };
}


export function SchemasUpdateMessageStatusPayloadToJSON(value?: SchemasUpdateMessageStatusPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'isRead': value.isRead,
    };
}
