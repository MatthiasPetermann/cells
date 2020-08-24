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

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _IdmRole = require('./IdmRole');

var _IdmRole2 = _interopRequireDefault(_IdmRole);

var _ServiceResourcePolicy = require('./ServiceResourcePolicy');

var _ServiceResourcePolicy2 = _interopRequireDefault(_ServiceResourcePolicy);

/**
* The IdmUser model module.
* @module model/IdmUser
* @version 1.0
*/

var IdmUser = (function () {
    /**
    * Constructs a new <code>IdmUser</code>.
    * @alias module:model/IdmUser
    * @class
    */

    function IdmUser() {
        _classCallCheck(this, IdmUser);

        this.Uuid = undefined;
        this.GroupPath = undefined;
        this.Attributes = undefined;
        this.Roles = undefined;
        this.Login = undefined;
        this.Password = undefined;
        this.OldPassword = undefined;
        this.IsGroup = undefined;
        this.GroupLabel = undefined;
        this.LastConnected = undefined;
        this.Policies = undefined;
        this.PoliciesContextEditable = undefined;
    }

    /**
    * Constructs a <code>IdmUser</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/IdmUser} obj Optional instance to populate.
    * @return {module:model/IdmUser} The populated <code>IdmUser</code> instance.
    */

    IdmUser.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IdmUser();

            if (data.hasOwnProperty('Uuid')) {
                obj['Uuid'] = _ApiClient2['default'].convertToType(data['Uuid'], 'String');
            }
            if (data.hasOwnProperty('GroupPath')) {
                obj['GroupPath'] = _ApiClient2['default'].convertToType(data['GroupPath'], 'String');
            }
            if (data.hasOwnProperty('Attributes')) {
                obj['Attributes'] = _ApiClient2['default'].convertToType(data['Attributes'], { 'String': 'String' });
            }
            if (data.hasOwnProperty('Roles')) {
                obj['Roles'] = _ApiClient2['default'].convertToType(data['Roles'], [_IdmRole2['default']]);
            }
            if (data.hasOwnProperty('Login')) {
                obj['Login'] = _ApiClient2['default'].convertToType(data['Login'], 'String');
            }
            if (data.hasOwnProperty('Password')) {
                obj['Password'] = _ApiClient2['default'].convertToType(data['Password'], 'String');
            }
            if (data.hasOwnProperty('OldPassword')) {
                obj['OldPassword'] = _ApiClient2['default'].convertToType(data['OldPassword'], 'String');
            }
            if (data.hasOwnProperty('IsGroup')) {
                obj['IsGroup'] = _ApiClient2['default'].convertToType(data['IsGroup'], 'Boolean');
            }
            if (data.hasOwnProperty('GroupLabel')) {
                obj['GroupLabel'] = _ApiClient2['default'].convertToType(data['GroupLabel'], 'String');
            }
            if (data.hasOwnProperty('LastConnected')) {
                obj['LastConnected'] = _ApiClient2['default'].convertToType(data['LastConnected'], 'Number');
            }
            if (data.hasOwnProperty('Policies')) {
                obj['Policies'] = _ApiClient2['default'].convertToType(data['Policies'], [_ServiceResourcePolicy2['default']]);
            }
            if (data.hasOwnProperty('PoliciesContextEditable')) {
                obj['PoliciesContextEditable'] = _ApiClient2['default'].convertToType(data['PoliciesContextEditable'], 'Boolean');
            }
        }
        return obj;
    };

    /**
    * @member {String} Uuid
    */
    return IdmUser;
})();

exports['default'] = IdmUser;
module.exports = exports['default'];

/**
* @member {String} GroupPath
*/

/**
* @member {Object.<String, String>} Attributes
*/

/**
* @member {Array.<module:model/IdmRole>} Roles
*/

/**
* @member {String} Login
*/

/**
* @member {String} Password
*/

/**
* @member {String} OldPassword
*/

/**
* @member {Boolean} IsGroup
*/

/**
* @member {String} GroupLabel
*/

/**
* @member {Number} LastConnected
*/

/**
* @member {Array.<module:model/ServiceResourcePolicy>} Policies
*/

/**
* Context-resolved to quickly check if user is editable or not.
* @member {Boolean} PoliciesContextEditable
*/
