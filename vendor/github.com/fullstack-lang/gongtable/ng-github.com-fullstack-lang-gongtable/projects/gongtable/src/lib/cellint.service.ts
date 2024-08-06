// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { CellIntAPI } from './cellint-api'
import { CellInt, CopyCellIntToCellIntAPI } from './cellint'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class CellIntService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CellIntServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private cellintsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.cellintsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/cellints';
  }

  /** GET cellints from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI[]> {
    return this.getCellInts(GONG__StackPath, frontRepo)
  }
  getCellInts(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<CellIntAPI[]>(this.cellintsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<CellIntAPI[]>('getCellInts', []))
      );
  }

  /** GET cellint by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI> {
    return this.getCellInt(id, GONG__StackPath, frontRepo)
  }
  getCellInt(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.cellintsUrl}/${id}`;
    return this.http.get<CellIntAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched cellint id=${id}`)),
      catchError(this.handleError<CellIntAPI>(`getCellInt id=${id}`))
    );
  }

  // postFront copy cellint to a version with encoded pointers and post to the back
  postFront(cellint: CellInt, GONG__StackPath: string): Observable<CellIntAPI> {
    let cellintAPI = new CellIntAPI
    CopyCellIntToCellIntAPI(cellint, cellintAPI)
    const id = typeof cellintAPI === 'number' ? cellintAPI : cellintAPI.ID
    const url = `${this.cellintsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CellIntAPI>(url, cellintAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CellIntAPI>('postCellInt'))
    );
  }
  
  /** POST: add a new cellint to the server */
  post(cellintdb: CellIntAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI> {
    return this.postCellInt(cellintdb, GONG__StackPath, frontRepo)
  }
  postCellInt(cellintdb: CellIntAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CellIntAPI>(this.cellintsUrl, cellintdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted cellintdb id=${cellintdb.ID}`)
      }),
      catchError(this.handleError<CellIntAPI>('postCellInt'))
    );
  }

  /** DELETE: delete the cellintdb from the server */
  delete(cellintdb: CellIntAPI | number, GONG__StackPath: string): Observable<CellIntAPI> {
    return this.deleteCellInt(cellintdb, GONG__StackPath)
  }
  deleteCellInt(cellintdb: CellIntAPI | number, GONG__StackPath: string): Observable<CellIntAPI> {
    const id = typeof cellintdb === 'number' ? cellintdb : cellintdb.ID;
    const url = `${this.cellintsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CellIntAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted cellintdb id=${id}`)),
      catchError(this.handleError<CellIntAPI>('deleteCellInt'))
    );
  }

  // updateFront copy cellint to a version with encoded pointers and update to the back
  updateFront(cellint: CellInt, GONG__StackPath: string): Observable<CellIntAPI> {
    let cellintAPI = new CellIntAPI
    CopyCellIntToCellIntAPI(cellint, cellintAPI)
    const id = typeof cellintAPI === 'number' ? cellintAPI : cellintAPI.ID
    const url = `${this.cellintsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<CellIntAPI>(url, cellintAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CellIntAPI>('updateCellInt'))
    );
  }

  /** PUT: update the cellintdb on the server */
  update(cellintdb: CellIntAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI> {
    return this.updateCellInt(cellintdb, GONG__StackPath, frontRepo)
  }
  updateCellInt(cellintdb: CellIntAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellIntAPI> {
    const id = typeof cellintdb === 'number' ? cellintdb : cellintdb.ID;
    const url = `${this.cellintsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CellIntAPI>(url, cellintdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated cellintdb id=${cellintdb.ID}`)
      }),
      catchError(this.handleError<CellIntAPI>('updateCellInt'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CellIntService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CellIntService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
