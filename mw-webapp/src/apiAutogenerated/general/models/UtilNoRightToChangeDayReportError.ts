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
 * @interface UtilNoRightToChangeDayReportError
 */
export interface UtilNoRightToChangeDayReportError {
    /**
     * 
     * @type {string}
     * @memberof UtilNoRightToChangeDayReportError
     */
    error: string;
    /**
     * 
     * @type {string}
     * @memberof UtilNoRightToChangeDayReportError
     */
    errorId: string;
}

/**
 * Check if a given object implements the UtilNoRightToChangeDayReportError interface.
 */
export function instanceOfUtilNoRightToChangeDayReportError(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "error" in value;
    isInstance = isInstance && "errorId" in value;

    return isInstance;
}

export function UtilNoRightToChangeDayReportErrorFromJSON(json: any): UtilNoRightToChangeDayReportError {
    return UtilNoRightToChangeDayReportErrorFromJSONTyped(json, false);
}

export function UtilNoRightToChangeDayReportErrorFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): UtilNoRightToChangeDayReportError {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'error': json['error'],
        'errorId': json['errorId'],
    };
}


export function UtilNoRightToChangeDayReportErrorToJSON(value?: UtilNoRightToChangeDayReportError | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'error': value.error,
        'errorId': value.errorId,
    };
}
