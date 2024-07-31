// @ts-nocheck
/* eslint-disable */
/**
 * Masters way general API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { SchemasWayPlainResponse } from './SchemasWayPlainResponse';
import {
    SchemasWayPlainResponseFromJSON,
    SchemasWayPlainResponseFromJSONTyped,
    SchemasWayPlainResponseToJSON,
} from './SchemasWayPlainResponse';

/**
 * 
 * @export
 * @interface SchemasWayCollectionPopulatedResponse
 */
export interface SchemasWayCollectionPopulatedResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasWayCollectionPopulatedResponse
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasWayCollectionPopulatedResponse
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasWayCollectionPopulatedResponse
     */
    ownerUuid: string;
    /**
     * should be removed after separation custom collections and default pseudocollections
     * @type {string}
     * @memberof SchemasWayCollectionPopulatedResponse
     */
    type: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasWayCollectionPopulatedResponse
     */
    updatedAt: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasWayCollectionPopulatedResponse
     */
    uuid: string;
    /**
     * 
     * @type {Array<SchemasWayPlainResponse>}
     * @memberof SchemasWayCollectionPopulatedResponse
     */
    ways: Array<SchemasWayPlainResponse>;
}

/**
 * Check if a given object implements the SchemasWayCollectionPopulatedResponse interface.
 */
export function instanceOfSchemasWayCollectionPopulatedResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "createdAt" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "ownerUuid" in value;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "updatedAt" in value;
    isInstance = isInstance && "uuid" in value;
    isInstance = isInstance && "ways" in value;

    return isInstance;
}

export function SchemasWayCollectionPopulatedResponseFromJSON(json: any): SchemasWayCollectionPopulatedResponse {
    return SchemasWayCollectionPopulatedResponseFromJSONTyped(json, false);
}

export function SchemasWayCollectionPopulatedResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasWayCollectionPopulatedResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'],
        'name': json['name'],
        'ownerUuid': json['ownerUuid'],
        'type': json['type'],
        'updatedAt': json['updatedAt'],
        'uuid': json['uuid'],
        'ways': ((json['ways'] as Array<any>).map(SchemasWayPlainResponseFromJSON)),
    };
}


export function SchemasWayCollectionPopulatedResponseToJSON(value?: SchemasWayCollectionPopulatedResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'createdAt': value.createdAt,
        'name': value.name,
        'ownerUuid': value.ownerUuid,
        'type': value.type,
        'updatedAt': value.updatedAt,
        'uuid': value.uuid,
        'ways': ((value.ways as Array<any>).map(SchemasWayPlainResponseToJSON)),
    };
}

