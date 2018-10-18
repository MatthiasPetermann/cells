/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from '../ApiClient';
import JobsAction from './JobsAction';
import JobsSchedule from './JobsSchedule';
import JobsTask from './JobsTask';





/**
* The JobsJob model module.
* @module model/JobsJob
* @version 1.0
*/
export default class JobsJob {
    /**
    * Constructs a new <code>JobsJob</code>.
    * @alias module:model/JobsJob
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>JobsJob</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/JobsJob} obj Optional instance to populate.
    * @return {module:model/JobsJob} The populated <code>JobsJob</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobsJob();

            
            
            

            if (data.hasOwnProperty('ID')) {
                obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('Label')) {
                obj['Label'] = ApiClient.convertToType(data['Label'], 'String');
            }
            if (data.hasOwnProperty('Owner')) {
                obj['Owner'] = ApiClient.convertToType(data['Owner'], 'String');
            }
            if (data.hasOwnProperty('Inactive')) {
                obj['Inactive'] = ApiClient.convertToType(data['Inactive'], 'Boolean');
            }
            if (data.hasOwnProperty('Languages')) {
                obj['Languages'] = ApiClient.convertToType(data['Languages'], ['String']);
            }
            if (data.hasOwnProperty('EventNames')) {
                obj['EventNames'] = ApiClient.convertToType(data['EventNames'], ['String']);
            }
            if (data.hasOwnProperty('Schedule')) {
                obj['Schedule'] = JobsSchedule.constructFromObject(data['Schedule']);
            }
            if (data.hasOwnProperty('AutoStart')) {
                obj['AutoStart'] = ApiClient.convertToType(data['AutoStart'], 'Boolean');
            }
            if (data.hasOwnProperty('AutoClean')) {
                obj['AutoClean'] = ApiClient.convertToType(data['AutoClean'], 'Boolean');
            }
            if (data.hasOwnProperty('Actions')) {
                obj['Actions'] = ApiClient.convertToType(data['Actions'], [JobsAction]);
            }
            if (data.hasOwnProperty('MaxConcurrency')) {
                obj['MaxConcurrency'] = ApiClient.convertToType(data['MaxConcurrency'], 'Number');
            }
            if (data.hasOwnProperty('TasksSilentUpdate')) {
                obj['TasksSilentUpdate'] = ApiClient.convertToType(data['TasksSilentUpdate'], 'Boolean');
            }
            if (data.hasOwnProperty('Tasks')) {
                obj['Tasks'] = ApiClient.convertToType(data['Tasks'], [JobsTask]);
            }
        }
        return obj;
    }

    /**
    * @member {String} ID
    */
    ID = undefined;
    /**
    * @member {String} Label
    */
    Label = undefined;
    /**
    * @member {String} Owner
    */
    Owner = undefined;
    /**
    * @member {Boolean} Inactive
    */
    Inactive = undefined;
    /**
    * @member {Array.<String>} Languages
    */
    Languages = undefined;
    /**
    * @member {Array.<String>} EventNames
    */
    EventNames = undefined;
    /**
    * @member {module:model/JobsSchedule} Schedule
    */
    Schedule = undefined;
    /**
    * @member {Boolean} AutoStart
    */
    AutoStart = undefined;
    /**
    * @member {Boolean} AutoClean
    */
    AutoClean = undefined;
    /**
    * @member {Array.<module:model/JobsAction>} Actions
    */
    Actions = undefined;
    /**
    * @member {Number} MaxConcurrency
    */
    MaxConcurrency = undefined;
    /**
    * @member {Boolean} TasksSilentUpdate
    */
    TasksSilentUpdate = undefined;
    /**
    * @member {Array.<module:model/JobsTask>} Tasks
    */
    Tasks = undefined;








}


