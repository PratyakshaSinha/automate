import {
  Component,
  EventEmitter,
  Input,
  OnDestroy,
  OnInit
} from '@angular/core';
import { Subject } from 'rxjs';
import { FormBuilder,  Validators, FormGroup } from '@angular/forms';
import {  takeUntil } from 'rxjs/operators';
import { Regex } from 'app/helpers/auth/regex';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { InfraRole } from 'app/entities/infra-roles/infra-role.model';

enum UrlTestState {
  Inactive,
  Loading,
  Success,
  Failure
}

@Component({
  selector: 'app-create-infra-role-modal',
  templateUrl: './create-infra-role-modal.component.html',
  styleUrls: ['./create-infra-role-modal.component.scss']
})
export class CreateInfraRoleModalComponent implements OnInit, OnDestroy {
  @Input() openEvent: EventEmitter<void>;
  @Input() rolesList: InfraRole[] = [];
  @Input() serverId: string;
  @Input() orgId: string;
  public visible = false;
  public creating = false;
  public sending = false;
  public conflictErrorEvent = new EventEmitter<boolean>();
  public close = new EventEmitter();
  public createForm: FormGroup;
  public hookStatus = UrlTestState.Inactive;
  public conflictError = false;
  public urlState = UrlTestState;
  public targetKeys: string[];
  public alertTypeKeys: string[];

  public isLinear = true;
  public firstFormGroup: FormGroup;
  public secondFormGroup: FormGroup;
  public thirdFormGroup: FormGroup;
  public fourthFormGroup: FormGroup;

  basket = [];

  baskets = [];
  public server: string;
  public org: string;
  public showdrag = false;
  items: InfraRole[] = [];
  private isDestroyed = new Subject<boolean>();
  public jsonString: string;
  public dattrParseError: boolean;
  public oattrParseError: boolean;
  public data: any;
  public textareaID: string;

  constructor(
    private fb: FormBuilder
  ) {

    this.firstFormGroup = this.fb.group({
      name: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
      description: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]]
    });
    this.secondFormGroup = this.fb.group({
      
    });
    this.thirdFormGroup = this.fb.group({
      dattr: [''],
    });
    this.fourthFormGroup = this.fb.group({
      oattr: ['',[Validators.required]]
    });
  }

  ngOnInit() {
    this.openEvent.pipe(takeUntil(this.isDestroyed))
      .subscribe(() => {
        this.conflictError = false;
        this.visible = true;
        this.items = this.rolesList;

      this.showdrag = true;
      this.server = this.serverId;
      this.org = this.orgId;

      });

  }

  ngOnDestroy(): void {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }


  drop(event: CdkDragDrop<string[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    } else {
      transferArrayItem(event.previousContainer.data,
                        event.container.data,
                        event.previousIndex,
                        event.currentIndex);
    }
  }

  public handleInput(event: KeyboardEvent): void {
    if (this.isNavigationKey(event)) {
      return;
    }
    this.conflictError = false;
  }

  closeCreateModal(): void {
    this.resetCreateModal();
    this.visible = false;
  }

  createNotification(): void {
    this.creating = true;
    
    const role = {
      id: this.createForm.controls['id'].value,
      name: this.createForm.controls['name'].value.trim(),
      description: this.createForm.controls['description'].value.trim(),
      dattr: this.createForm.controls['dattr'].value.trim()
    };
    console.log(role);

  }

  private resetCreateModal(): void {
    this.hookStatus = UrlTestState.Inactive;
    this.creating = false;
    this.firstFormGroup.reset();
    this.secondFormGroup.reset();
    this.conflictErrorEvent.emit(false);
  }

  private isNavigationKey(event: KeyboardEvent): boolean {
    return event.key === 'Shift' || event.key === 'Tab';
  }

  // this is the initial value set to the component
  public writeValue(obj: any) {
    if (obj) {
        this.data = obj;
        // this will format it with 4 character spacing
        this.jsonString = JSON.stringify(this.data, undefined, 4); 
    }
  }

  // registers 'fn' that will be fired wheb changes are made
  // this is how we emit the changes back to the form
  public registerOnChange(fn: any) {
      this.propagateChange = fn;
  }

  private onChangeJSON(event) {
    this.dattrParseError = false;
    this.oattrParseError = false;
    // get value from text area
    let newValue = event.target.value;
    this.textareaID = event.target.id;
    try {
        // parse it to json
        this.data = JSON.parse(newValue);
        this.textareaID == 'dattr' ? this.dattrParseError = false : '';
        this.textareaID == 'oattr' ? this.oattrParseError = false : '';
    } catch (ex) {
        // set parse error if it fails
        this.textareaID == 'dattr' ? this.dattrParseError = true : '';
        this.textareaID == 'oattr' ? this.oattrParseError = true : '';
    }
    // update the form
    this.propagateChange(this.data);
  }

  // the method set in registerOnChange to emit changes back to the form
  private propagateChange = (_: any) => { };

}
