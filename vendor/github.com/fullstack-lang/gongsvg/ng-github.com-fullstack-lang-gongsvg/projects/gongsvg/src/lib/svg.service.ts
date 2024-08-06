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

import { SVGAPI } from './svg-api'
import { SVG, CopySVGToSVGAPI } from './svg'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { LayerAPI } from './layer-api'
import { RectAPI } from './rect-api'

@Injectable({
  providedIn: 'root'
})
export class SVGService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  SVGServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private svgsUrl: string

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
    this.svgsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/svgs';
  }

  /** GET svgs from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI[]> {
    return this.getSVGs(GONG__StackPath, frontRepo)
  }
  getSVGs(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<SVGAPI[]>(this.svgsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<SVGAPI[]>('getSVGs', []))
      );
  }

  /** GET svg by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI> {
    return this.getSVG(id, GONG__StackPath, frontRepo)
  }
  getSVG(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.svgsUrl}/${id}`;
    return this.http.get<SVGAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched svg id=${id}`)),
      catchError(this.handleError<SVGAPI>(`getSVG id=${id}`))
    );
  }

  // postFront copy svg to a version with encoded pointers and post to the back
  postFront(svg: SVG, GONG__StackPath: string): Observable<SVGAPI> {
    let svgAPI = new SVGAPI
    CopySVGToSVGAPI(svg, svgAPI)
    const id = typeof svgAPI === 'number' ? svgAPI : svgAPI.ID
    const url = `${this.svgsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SVGAPI>(url, svgAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SVGAPI>('postSVG'))
    );
  }
  
  /** POST: add a new svg to the server */
  post(svgdb: SVGAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI> {
    return this.postSVG(svgdb, GONG__StackPath, frontRepo)
  }
  postSVG(svgdb: SVGAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SVGAPI>(this.svgsUrl, svgdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted svgdb id=${svgdb.ID}`)
      }),
      catchError(this.handleError<SVGAPI>('postSVG'))
    );
  }

  /** DELETE: delete the svgdb from the server */
  delete(svgdb: SVGAPI | number, GONG__StackPath: string): Observable<SVGAPI> {
    return this.deleteSVG(svgdb, GONG__StackPath)
  }
  deleteSVG(svgdb: SVGAPI | number, GONG__StackPath: string): Observable<SVGAPI> {
    const id = typeof svgdb === 'number' ? svgdb : svgdb.ID;
    const url = `${this.svgsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<SVGAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted svgdb id=${id}`)),
      catchError(this.handleError<SVGAPI>('deleteSVG'))
    );
  }

  // updateFront copy svg to a version with encoded pointers and update to the back
  updateFront(svg: SVG, GONG__StackPath: string): Observable<SVGAPI> {
    let svgAPI = new SVGAPI
    CopySVGToSVGAPI(svg, svgAPI)
    const id = typeof svgAPI === 'number' ? svgAPI : svgAPI.ID
    const url = `${this.svgsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<SVGAPI>(url, svgAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SVGAPI>('updateSVG'))
    );
  }

  /** PUT: update the svgdb on the server */
  update(svgdb: SVGAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI> {
    return this.updateSVG(svgdb, GONG__StackPath, frontRepo)
  }
  updateSVG(svgdb: SVGAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGAPI> {
    const id = typeof svgdb === 'number' ? svgdb : svgdb.ID;
    const url = `${this.svgsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<SVGAPI>(url, svgdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated svgdb id=${svgdb.ID}`)
      }),
      catchError(this.handleError<SVGAPI>('updateSVG'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in SVGService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("SVGService" + error); // log to console instead

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
