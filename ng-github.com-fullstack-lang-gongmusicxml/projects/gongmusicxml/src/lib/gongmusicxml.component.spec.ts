import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmusicxmlComponent } from './gongmusicxml.component';

describe('GongmusicxmlComponent', () => {
  let component: GongmusicxmlComponent;
  let fixture: ComponentFixture<GongmusicxmlComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongmusicxmlComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongmusicxmlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
