package templates

var DtoTemplate = `{{with $pMN := formatModuleName $.ModuleName }}/*
 * CLIENT DTOS
 */
export interface ICreate{{$pMN}}ClientDTO extends I{{$pMN}}ClientDTO {}
export interface I{{$pMN}}ClientDTO extends I{{$pMN}}ServiceDTO {}


/*
 * MESSAGE DTOS
 */
export interface ICreate{{$pMN}}MessageDTO extends I{{$pMN}}MessageDTO {}
export interface I{{$pMN}}MessageDTO extends I{{$pMN}}ServiceDTO {}


/*
 * SERVICE DTOS
 */
export interface ICreate{{$pMN}}ServiceDTO extends I{{$pMN}}ServiceDTO {}
export interface I{{$pMN}}ServiceDTO extends I{{$pMN}}DataDTO {}


/*
 * DATA DTOS
 */
export interface ICreate{{$pMN}}DataDTO extends I{{$pMN}}DataDTO {}
export interface I{{$pMN}}DataDTO {}
{{end}}`
