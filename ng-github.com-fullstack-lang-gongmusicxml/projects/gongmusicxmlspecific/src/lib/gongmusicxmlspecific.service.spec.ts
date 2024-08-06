import { TestBed } from '@angular/core/testing';

import { GongmusicxmlspecificService } from './gongmusicxmlspecific.service';

describe('GongmusicxmlspecificService', () => {
  let service: GongmusicxmlspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongmusicxmlspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
