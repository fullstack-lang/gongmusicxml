import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmusicxmlspecificComponent } from './gongmusicxmlspecific.component';

describe('GongmusicxmlspecificComponent', () => {
  let component: GongmusicxmlspecificComponent;
  let fixture: ComponentFixture<GongmusicxmlspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongmusicxmlspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongmusicxmlspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
