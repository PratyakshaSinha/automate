import { HttpErrorResponse } from '@angular/common/http';
import { Action } from '@ngrx/store';
import { InfraRole } from './infra-role.model';

export enum RoleActionTypes {
  GET_ALL = 'ROLES::GET_ALL',
  GET_ALL_SUCCESS = 'ROLES::GET_ALL::SUCCESS',
  GET_ALL_FAILURE = 'ROLES::GET_ALL::FAILURE',
  GET = 'ROLES::GET',
  GET_SUCCESS = 'ROLES::GET::SUCCESS',
  GET_FAILURE = 'ROLES::GET::FAILURE',
  CREATE            = 'ROLES::CREATE',
  CREATE_SUCCESS    = 'ROLES::CREATE::SUCCESS',
  CREATE_FAILURE    = 'ROLES::CREATE::FAILURE',
}

export interface RolesSuccessPayload {
  roles: InfraRole[];
}

export interface RoleSuccessPayload {
  role: InfraRole;
}

export class GetRoles implements Action {
  readonly type = RoleActionTypes.GET_ALL;
  constructor(public payload: { server_id: string, org_id: string }) { }
}

export class GetRolesSuccess implements Action {
  readonly type = RoleActionTypes.GET_ALL_SUCCESS;
  constructor(public payload: RolesSuccessPayload) { }
}

export class GetRolesFailure implements Action {
  readonly type = RoleActionTypes.GET_ALL_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class GetRole implements Action {
  readonly type = RoleActionTypes.GET;
  constructor(public payload: { server_id: string, org_id: string, name: string }) { }
}

export class GetRoleSuccess implements Action {
  readonly type = RoleActionTypes.GET_SUCCESS;
  constructor(public payload: InfraRole) { }
}

export class GetRoleFailure implements Action {
  readonly type = RoleActionTypes.GET_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export interface CreateRolePayload {
  org_id: string;
  server_id: string;
  name: string;
  description: string;
  default_attributes: Object ;
  override_attributes: Object ;
  run_list: string[];
  env_run_lists: Object;
}

export class CreateRole implements Action {
  readonly type = RoleActionTypes.CREATE;
  constructor(public payload: { server_id: string, org_id: string, role: CreateRolePayload } ) { }
}

export class CreateRoleSuccess implements Action {
  readonly type = RoleActionTypes.CREATE_SUCCESS;
  constructor(public payload) { }
}

export class CreateRoleFailure implements Action {
  readonly type = RoleActionTypes.CREATE_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export type RoleActions =
  | GetRoles
  | GetRolesSuccess
  | GetRolesFailure
  | GetRole
  | GetRoleSuccess
  | GetRoleFailure
  | CreateRole
  | CreateRoleSuccess
  | CreateRoleFailure;
