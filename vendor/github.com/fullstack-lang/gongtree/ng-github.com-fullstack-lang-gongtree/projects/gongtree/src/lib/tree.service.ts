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

import { TreeAPI } from './tree-api'
import { Tree, CopyTreeToTreeAPI } from './tree'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { NodeAPI } from './node-api'

@Injectable({
  providedIn: 'root'
})
export class TreeService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  TreeServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private treesUrl: string

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
    this.treesUrl = origin + '/api/github.com/fullstack-lang/gongtree/go/v1/trees';
  }

  /** GET trees from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI[]> {
    return this.getTrees(GONG__StackPath, frontRepo)
  }
  getTrees(GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<TreeAPI[]>(this.treesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<TreeAPI[]>('getTrees', []))
      );
  }

  /** GET tree by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI> {
    return this.getTree(id, GONG__StackPath, frontRepo)
  }
  getTree(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.treesUrl}/${id}`;
    return this.http.get<TreeAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched tree id=${id}`)),
      catchError(this.handleError<TreeAPI>(`getTree id=${id}`))
    );
  }

  // postFront copy tree to a version with encoded pointers and post to the back
  postFront(tree: Tree, GONG__StackPath: string): Observable<TreeAPI> {
    let treeAPI = new TreeAPI
    CopyTreeToTreeAPI(tree, treeAPI)
    const id = typeof treeAPI === 'number' ? treeAPI : treeAPI.ID
    const url = `${this.treesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<TreeAPI>(url, treeAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<TreeAPI>('postTree'))
    );
  }
  
  /** POST: add a new tree to the server */
  post(treedb: TreeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI> {
    return this.postTree(treedb, GONG__StackPath, frontRepo)
  }
  postTree(treedb: TreeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<TreeAPI>(this.treesUrl, treedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted treedb id=${treedb.ID}`)
      }),
      catchError(this.handleError<TreeAPI>('postTree'))
    );
  }

  /** DELETE: delete the treedb from the server */
  delete(treedb: TreeAPI | number, GONG__StackPath: string): Observable<TreeAPI> {
    return this.deleteTree(treedb, GONG__StackPath)
  }
  deleteTree(treedb: TreeAPI | number, GONG__StackPath: string): Observable<TreeAPI> {
    const id = typeof treedb === 'number' ? treedb : treedb.ID;
    const url = `${this.treesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<TreeAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted treedb id=${id}`)),
      catchError(this.handleError<TreeAPI>('deleteTree'))
    );
  }

  // updateFront copy tree to a version with encoded pointers and update to the back
  updateFront(tree: Tree, GONG__StackPath: string): Observable<TreeAPI> {
    let treeAPI = new TreeAPI
    CopyTreeToTreeAPI(tree, treeAPI)
    const id = typeof treeAPI === 'number' ? treeAPI : treeAPI.ID
    const url = `${this.treesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<TreeAPI>(url, treeAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<TreeAPI>('updateTree'))
    );
  }

  /** PUT: update the treedb on the server */
  update(treedb: TreeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI> {
    return this.updateTree(treedb, GONG__StackPath, frontRepo)
  }
  updateTree(treedb: TreeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<TreeAPI> {
    const id = typeof treedb === 'number' ? treedb : treedb.ID;
    const url = `${this.treesUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<TreeAPI>(url, treedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated treedb id=${treedb.ID}`)
      }),
      catchError(this.handleError<TreeAPI>('updateTree'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in TreeService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("TreeService" + error); // log to console instead

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
