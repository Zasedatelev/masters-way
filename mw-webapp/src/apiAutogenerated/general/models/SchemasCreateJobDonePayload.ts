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
/**
 * 
 * @export
 * @interface SchemasCreateJobDonePayload
 */
export interface SchemasCreateJobDonePayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateJobDonePayload
     */
    dayReportUuid: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateJobDonePayload
     */
    description: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SchemasCreateJobDonePayload
     */
    jobTagUuids: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateJobDonePayload
     */
    ownerUuid: string;
    /**
     * 
     * @type {number}
     * @memberof SchemasCreateJobDonePayload
     */
    time: number;
}

/**
 * Check if a given object implements the SchemasCreateJobDonePayload interface.
 */
export function instanceOfSchemasCreateJobDonePayload(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "dayReportUuid" in value;
    isInstance = isInstance && "description" in value;
    isInstance = isInstance && "jobTagUuids" in value;
    isInstance = isInstance && "ownerUuid" in value;
    isInstance = isInstance && "time" in value;

    return isInstance;
}

export function SchemasCreateJobDonePayloadFromJSON(json: any): SchemasCreateJobDonePayload {
    return SchemasCreateJobDonePayloadFromJSONTyped(json, false);
}

export function SchemasCreateJobDonePayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasCreateJobDonePayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dayReportUuid': json['dayReportUuid'],
        'description': json['description'],
        'jobTagUuids': json['jobTagUuids'],
        'ownerUuid': json['ownerUuid'],
        'time': json['time'],
    };
}


export function SchemasCreateJobDonePayloadToJSON(value?: SchemasCreateJobDonePayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dayReportUuid': value.dayReportUuid,
        'description': value.description,
        'jobTagUuids': value.jobTagUuids,
        'ownerUuid': value.ownerUuid,
        'time': value.time,
    };
}

