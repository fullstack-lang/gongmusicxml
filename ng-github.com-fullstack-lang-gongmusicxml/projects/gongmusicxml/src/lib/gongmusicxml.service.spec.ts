import { TestBed } from '@angular/core/testing';

import { GongmusicxmlService } from './gongmusicxml.service';

describe('GongmusicxmlService', () => {
  let service: GongmusicxmlService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongmusicxmlService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
